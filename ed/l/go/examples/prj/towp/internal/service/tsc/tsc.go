package tsc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/common"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
)

type ServiceCatalog interface {
	InStorePickingEnabled(context.Context, string, string) (bool, error)
}

var _ ServiceCatalog = &Service{}

// LocationInfo holds config values about TSC location.
type LocationInfo struct {
	MfcID    string `json:"mfc-ref-code"`
	Type     string `json:"location-type"`
	Timezone string `json:"timezone"`
}

type Service struct {
	cfg        *config.Config
	logger     *zap.SugaredLogger
	httpClient *http.Client
}

func New(cfg *config.Config, logger *zap.SugaredLogger, httpClient *http.Client) *Service {
	return &Service{
		cfg:        cfg,
		logger:     logger,
		httpClient: httpClient,
	}
}

func (s *Service) InStorePickingEnabled(ctx context.Context, retailerID, mfcID string) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "tsc.isps-enabled")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		fmt.Sprintf(s.cfg.TSCURLTemplate+"/api/v1/configuration/config-items",
			retailerID, common.GetCtxEnv(ctx)),
		nil)
	if err != nil {
		return false, err
	}

	url := req.URL.Query()
	url.Add("location-codes", mfcID)
	url.Add("level", "mfc")
	url.Add("categories", s.cfg.ISPSCategory)
	url.Add("value-format", "json")

	req.URL.RawQuery = url.Encode()

	req.Header.Add("X-Token", common.GetCtxToken(ctx))

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300

	if !statusOK {
		return false, fmt.Errorf(
			"an error occurred while retrieving service catalog config, response: %s", string(rawBody))
	}

	responsePayload := make([]map[string]any, 0)
	if err := json.Unmarshal(rawBody, &responsePayload); err != nil {
		return false, fmt.Errorf("unable to unmarshal service-catalog config")
	}

	for _, item := range responsePayload {
		if item["name"] == s.cfg.ISPSEnabledConfigName {
			boolVal, ok := item["value"].(bool)
			if !ok {
				return false, fmt.Errorf("unable to convert IspsEnabledConfig config to bool, val: %#v", item["value"])
			}
			return boolVal, nil
		}
	}

	return false, fmt.Errorf("%s config is not found in service-catalog", s.cfg.ISPSEnabledConfigName)
}

// GetLocationInfo gets TSC location config.
func (s *Service) GetLocationInfo(ctx context.Context, retailerID, mfcID string) (LocationInfo, error) {
	ctx, span := trace.StartSpan(ctx, "tsc.get-location-info")
	defer span.End()

	env := common.GetCtxEnv(ctx)
	xToken := common.GetCtxToken(ctx)

	uri := fmt.Sprintf(s.cfg.TSCURLTemplate+"/api/v1/locations", retailerID, env)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("failed to create new request, err: %w", err)
	}
	req.Header.Add("X-Token", xToken)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("failed to perform request, err: %w", err)
	}
	defer res.Body.Close()

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("failed to read response body, err: %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return LocationInfo{}, fmt.Errorf("got unexpected http status code: %v", res.StatusCode)
	}
	locationsInfo := make([]LocationInfo, 0)
	err = json.Unmarshal(rawBody, &locationsInfo)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("failed to unmarshal response, err: %w", err)
	}

	for _, item := range locationsInfo {
		if mfcID == item.MfcID {
			return item, nil
		}
	}

	return LocationInfo{}, fmt.Errorf("location info not found")
}
