package mfcconfig

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/app/storage"
	"github.com/to-com/poc-td/app/tsc"
)

// MFCConfig describes MFCConfig service.
type MFCConfig interface {
	// List gets MFC configs.
	List(ctx context.Context, input payload.GetConfigsInput) (c payload.GetConfigsOutput, err error)
	// Get gets MFC config.
	Get(ctx context.Context) (c payload.MFCConfig, err error)
	// Update updates MFCConfig.
	Update(ctx context.Context, input payload.MFCConfig) (c payload.MFCConfig, err error)
	// Shutdown shuts down MFCConfig (stops config autorefresh).
	Shutdown(ctx context.Context)
}

// Service represents service to work with MFC configs.
type Service struct {
	storage         storage.Storage
	tsc             tsc.toServiceCatalog
	log             *zap.SugaredLogger
	autoRefreshTime int

	chanAutoRefreshDone chan bool
	cache               map[string]payload.MFCConfig
}

// New creates new MFCConfig instance.
func New(
	storage storage.Storage, tsc tsc.toServiceCatalog, log *zap.SugaredLogger, autoRefreshTime int,
) *Service {
	s := &Service{
		storage:         storage,
		tsc:             tsc,
		log:             log,
		autoRefreshTime: autoRefreshTime,
	}

	s.cache = make(map[string]payload.MFCConfig)
	s.refreshCache()

	if autoRefreshTime > 0 {
		s.startAutoRefresh(autoRefreshTime)
	}

	return s
}

// List {@inheritdoc}.
func (s *Service) List(ctx context.Context, input payload.GetConfigsInput) (c payload.GetConfigsOutput, err error) {
	configs, err := s.storage.ListMFCConfigs(ctx, input)
	if err != nil {
		return c, fmt.Errorf("failed to get MFCConfig, err: %w", err)
	}
	c.Configs = configs

	return c, nil
}

// Get {@inheritdoc}.
func (s *Service) Get(ctx context.Context) (c payload.MFCConfig, err error) {
	token := ctx.Value(payload.ContextKeyToken).(string)
	clientID := ctx.Value(payload.ContextKeyRetailer).(string)
	env := ctx.Value(payload.ContextKeyEnv).(string)
	mfcID := ctx.Value(payload.ContextKeyMfc).(string)

	// Get config from cache.
	cachedConfig, ok := s.getConfigFromCache(clientID, env, mfcID)
	if ok && !s.IsConfigExpired(cachedConfig) {
		return cachedConfig, nil
	}

	// Get config from DB.
	getFromDBInput := payload.GetConfigsInput{
		ClientID: clientID,
		Env:      env,
		MfcID:    mfcID,
	}
	dbConfigs, err := s.storage.ListMFCConfigs(ctx, getFromDBInput)
	if err != nil {
		return c, fmt.Errorf("failed to list MFC configs, err: %w", err)
	}
	if len(dbConfigs) > 1 {
		return c, fmt.Errorf("got more that 1 config for: %v", getFromDBInput)
	}
	if len(dbConfigs) > 0 && !s.IsConfigExpired(dbConfigs[0]) {
		s.putConfigIntoCache(dbConfigs[0])
		return dbConfigs[0], nil
	}

	// Get config from TSC.
	getFromTSCInput := payload.GetConfigInput{
		Token:    token,
		Retailer: clientID,
		Env:      env,
		MFC:      mfcID,
	}
	tscConfig, err := s.tsc.GetConfig(ctx, getFromTSCInput)
	if err != nil {
		return c, fmt.Errorf("failed to get OSR config %v, err: %w", getFromDBInput, err)
	}
	// Update config in DB.
	tscConfig.UpdatedAt = time.Now().UnixMilli()
	err = s.storage.UpdateMFCConfig(ctx, tscConfig)
	if err != nil {
		return c, fmt.Errorf("failed to update MFCConfig %v, err: %w", getFromDBInput, err)
	}
	// Update config in cache.
	s.putConfigIntoCache(tscConfig)

	return tscConfig, nil
}

// Update {@inheritdoc}.
func (s *Service) Update(ctx context.Context, input payload.MFCConfig) (c payload.MFCConfig, err error) {
	input.UpdatedAt = time.Now().UnixMilli()

	err = s.storage.UpdateMFCConfig(ctx, input)
	if err != nil {
		return c, fmt.Errorf("failed to update MFCConfig %v, err: %w", c, err)
	}

	s.putConfigIntoCache(input)

	return input, nil
}

// IsConfigExpired returns true in case when config updated earlier than now-autoRefreshTime.
func (s *Service) IsConfigExpired(conf payload.MFCConfig) bool {
	if s.autoRefreshTime == 0 {
		return false
	}

	threshold := time.Now().Add(-time.Duration(s.autoRefreshTime) * time.Second)
	updatedAt := time.UnixMilli(conf.UpdatedAt)

	return updatedAt.Before(threshold)
}

// Shutdown {@inheritdoc}.
func (s *Service) Shutdown(ctx context.Context) {
	s.stopAutoRefresh()
}

func (s *Service) startAutoRefresh(autoRefreshTime int) {
	s.chanAutoRefreshDone = make(chan bool)
	ticker := time.NewTicker(time.Duration(autoRefreshTime) * time.Second)

	go func() {
		for {
			select {
			case <-s.chanAutoRefreshDone:
				ticker.Stop()
				return
			case <-ticker.C:
				s.refreshCache()
			}
		}
	}()
}

func getCacheKey(clientID, env, mfcID string) string {
	return clientID + env + mfcID
}

func (s *Service) putConfigIntoCache(conf payload.MFCConfig) {
	key := getCacheKey(conf.ClientID, conf.Env, conf.MfcID)
	s.cache[key] = conf
}

func (s *Service) getConfigFromCache(clientID, env, mfcID string) (conf payload.MFCConfig, ok bool) {
	key := getCacheKey(clientID, env, mfcID)
	conf, ok = s.cache[key]

	return conf, ok
}

func (s *Service) refreshCache() {
	ctx := context.Background()
	input := payload.GetConfigsInput{}
	configs, err := s.storage.ListMFCConfigs(ctx, input)
	if err != nil {
		s.log.Errorf("failed to perform storage.ListMFCConfigs, err: %v", err)
		return
	}

	for _, c := range configs {
		s.putConfigIntoCache(c)
	}
}

func (s *Service) stopAutoRefresh() {
	if s.chanAutoRefreshDone == nil {
		return
	}

	s.chanAutoRefreshDone <- true
}
