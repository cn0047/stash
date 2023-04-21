//go:generate mockgen -source business.go -destination mock/mocks_for_business.go -package mockbusiness
package business

import (
	"context"
	"errors"
	"fmt"
	"github.com/to-com/wp/internal/business/validator"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/common"
	"github.com/to-com/wp/internal/dto"
	r "github.com/to-com/wp/internal/repository"
	"github.com/to-com/wp/internal/service/tsc"
)

var (
	// ErrTriggerTimeExceedsUpperBound represents error when time for trigger
	// is higher than time now + TriggersGenerationHoursLimit.
	ErrTriggerTimeExceedsUpperBound = errors.New("trigger time exceeds upper bound")
)

type DataStore interface {
	Createwp(ctx context.Context, wp r.wp) (r.wp, error)
	Getwp(ctx context.Context, retailerID, mfcID string) ([]r.wpItem, error)
	CreateTriggers(ctx context.Context, triggers []r.Trigger) (int, error)
	GetScheduleItems(ctx context.Context) ([]r.ScheduleItem, error)
	GetTriggers(ctx context.Context, retailerID, mfcID string) ([]r.TriggerItem, error)
	SelectTriggersToFire(ctx context.Context) ([]r.TriggerItem, error)
	MarkTriggersAsFired(ctx context.Context, triggers []r.TriggerItem, firedAt time.Time) error
}

type Publisher interface {
	PublishMessage(ctx context.Context, topicID string, message any, attrs map[string]string) (*string, error)
}

type Business struct {
	cfg            *config.Config
	logger         *zap.SugaredLogger
	repository     DataStore
	pubsub         Publisher
	serviceCatalog *tsc.Service
}

func New(
	cfg *config.Config, logger *zap.SugaredLogger, repository DataStore, pubsub Publisher,
	serviceCatalog *tsc.Service,
) *Business {
	return &Business{
		cfg:            cfg,
		logger:         logger,
		repository:     repository,
		pubsub:         pubsub,
		serviceCatalog: serviceCatalog,
	}
}

func (b *Business) Createwp(ctx context.Context, wp dto.wpRequest) (dto.wpResponse, error) {
	ispsEnabled, apiErr := b.serviceCatalog.InStorePickingEnabled(ctx, wp.RetailerID, wp.MfcID)
	if apiErr != nil {
		return dto.wpResponse{}, apiErr
	}

	var opts []validator.Options
	if ispsEnabled {
		opts = append(opts, validator.WithISPS())
	}
	v := validator.New(opts...)
	if validationErr := v.Validate(wp); validationErr != nil {
		return dto.wpResponse{}, validationErr
	}

	locationInfo, apiErr := b.serviceCatalog.GetLocationInfo(ctx, wp.RetailerID, wp.MfcID)
	if apiErr != nil {
		return dto.wpResponse{}, fmt.Errorf("failed to get timezone, err: %w", apiErr)
	}

	wpSpannerView := r.wpToSpannerView(wp)
	wpSpannerView.Timezone = locationInfo.Timezone
	wp, err := b.repository.Createwp(ctx, wpSpannerView)
	if err != nil {
		return dto.wpResponse{}, err
	}

	wpResponse := r.wpFromSpannerToResponseView(wp)

	pubsubAttributes := map[string]string{
		"env_type":    common.GetCtxEnv(ctx),
		"event_type":  "wp.Created",
		"retailer_id": wp.RetailerID,
		"mfc_id":      wp.MfcID,
		"source":      "wp",
	}

	message := dto.wpCreatedEvent{
		CreatedAt: time.Now(),
		wp:  wpResponse,
	}

	_, err = b.pubsub.PublishMessage(ctx, b.cfg.wpTopic, message, pubsubAttributes)
	if err != nil {
		b.logger.Errorf("Failed to publish wave plan: %s, error: %v", wp.ID, err)
	} else {
		b.logger.Infof("wp %s for Retailer: %s, Mfc: %s successfully published",
			wpResponse.ID, wpResponse.RetailerID, wpResponse.MfcID)
	}

	// Generate triggers for just created wp.
	scheduleItems := make([]r.ScheduleItem, 0)
	for _, wave := range wp.Waves {
		for _, schedule := range wave.Schedules {
			scheduleItem := r.ScheduleItem{
				RetailerID:   wp.RetailerID,
				MfcID:        wp.MfcID,
				Timezone:     wp.Timezone,
				wpID:   wp.ID,
				WaveID:       wave.ID,
				ScheduleID:   schedule.ID,
				ScheduleTime: schedule.ScheduleTime,
				CutoffTime:   wave.Cutoff,
			}
			scheduleItems = append(scheduleItems, scheduleItem)
		}
	}
	triggers, err := b.generateTriggersForScheduleItems(scheduleItems)
	if err != nil {
		b.logger.Errorf("error occurred while generating triggers: %v", err)
	}

	_, err = b.createTriggers(ctx, triggers)
	if err != nil {
		b.logger.Errorf("error occurred while creating triggers: %v", err)
	}

	return wpResponse, nil
}

func (b *Business) Getwp(ctx context.Context, retailerID, mfcID string) (dto.wpResponse, error) {
	wpItems, err := b.repository.Getwp(ctx, retailerID, mfcID)
	if err != nil {
		b.logger.Errorf("failed to get wave plan for retailer: %s, MFC: %s, err: %+v", retailerID, mfcID, err)
		return dto.wpResponse{}, err
	}

	if len(wpItems) == 0 {
		return dto.wpResponse{}, nil
	}

	return r.wpItemsFromSpannerToResponseView(wpItems), nil
}

func (b *Business) createTriggers(ctx context.Context, triggers []r.Trigger) (dto.GenerateTriggersResponse, error) {
	if len(triggers) == 0 {
		return dto.GenerateTriggersResponse{}, nil
	}
	insertedCount, err := b.repository.CreateTriggers(ctx, triggers)
	if err != nil {
		b.logger.Errorf("failed to save triggers to db: %v", err)
		return dto.GenerateTriggersResponse{}, err
	}
	return dto.GenerateTriggersResponse{
		GeneratedTriggers: insertedCount,
	}, nil
}

func (b *Business) getScheduleItems(ctx context.Context) ([]r.ScheduleItem, error) {
	scheduleItems, err := b.repository.GetScheduleItems(ctx)
	if err != nil {
		b.logger.Error("error occurred while getting schedules from db")
		return make([]r.ScheduleItem, 0), err
	}
	return scheduleItems, nil
}

// GenerateTriggers generates triggers for all retailers and all MFCs.
func (b *Business) GenerateTriggers(ctx context.Context) (res dto.GenerateTriggersResponse, err error) {
	scheduleItems, err := b.getScheduleItems(ctx)
	if err != nil {
		b.logger.Errorf("failed to get schedule items: %v", err)
		return dto.GenerateTriggersResponse{}, err
	}
	if len(scheduleItems) == 0 {
		return res, nil
	}
	triggers, err := b.generateTriggersForScheduleItems(scheduleItems)
	if err != nil {
		return res, err
	}
	return b.createTriggers(ctx, triggers)
}

func (b *Business) generateTriggersForScheduleItems(scheduleItems []r.ScheduleItem,
) (res []r.Trigger, err error) {
	var triggers []r.Trigger
	for _, scheduleItem := range scheduleItems {
		trigger, err := generateTriggerForScheduleItem(b.cfg.TriggersGenerationHoursLimit, scheduleItem)
		if err != nil {
			if errors.Is(err, ErrTriggerTimeExceedsUpperBound) {
				b.logger.Debugf("skipping generate triggers for schedule: %v, reason: %v", scheduleItem.ScheduleID, err)
				continue
			}
			return res, fmt.Errorf("failed to generate triggers for schedule: %v, err: %w", scheduleItem.ScheduleID, err)
		}
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

// generateTriggerForScheduleItem generates trigger for particular ScheduleItem.
func generateTriggerForScheduleItem(
	generationHoursLimit int, scheduleItem r.ScheduleItem,
) (res r.Trigger, err error) {
	location, err := time.LoadLocation(scheduleItem.Timezone)
	if err != nil {
		return res, fmt.Errorf("failed to load location for: %s, err: %w", scheduleItem.Timezone, err)
	}

	// IMPORTANT: triggerAt holds time in provided timezone.
	triggerAt, err := getTriggerTime(scheduleItem.ScheduleTime, location, generationHoursLimit)
	if err != nil {
		return res, fmt.Errorf("failed to get trigger time, err: %w", err)
	}

	cutoffDateTime, err := convertTimeToDateTime(scheduleItem.CutoffTime, location)
	if err != nil {
		return res, fmt.Errorf("failed to get cutoff date time, err: %w", err)
	}

	if cutoffDateTime.Before(triggerAt) {
		cutoffDateTime = cutoffDateTime.Add(24 * time.Hour)
	}

	input := r.Trigger{
		wpID:     scheduleItem.wpID,
		WaveID:         scheduleItem.WaveID,
		ScheduleID:     scheduleItem.ScheduleID,
		TriggerAt:      triggerAt.UTC(), // IMPORTANT: time in UTC.
		CutoffDateTime: cutoffDateTime.UTC(),
	}

	return input, nil
}

func (b *Business) GetTriggers(ctx context.Context, retailerID, mfcID string) (dto.GetTriggersResponse, error) {
	triggers, err := b.repository.GetTriggers(ctx, retailerID, mfcID)
	if err != nil {
		b.logger.Errorf("failed to get triggers for retailer: %s, MFC: %s, err: %+v", retailerID, mfcID, err)
		return dto.GetTriggersResponse{}, err
	}
	return r.TriggersFromSpannerToResponseView(triggers), err
}

// getTriggerTime returns triggerTime in given timezone with respect time now and allowedHoursWindow bound.
func getTriggerTime(
	hourMinute string, location *time.Location, allowedHoursWindow int,
) (r time.Time, err error) {
	now := time.Now().In(location)

	triggerTime, err := convertTimeToDateTime(hourMinute, location)
	if err != nil {
		return r, err
	}
	upperBound := now.Add(time.Duration(allowedHoursWindow) * time.Hour)
	if triggerTime.After(upperBound) {
		return r, ErrTriggerTimeExceedsUpperBound
	}

	return triggerTime, nil
}

func convertTimeToDateTime(hourMinute string, location *time.Location) (dateTime time.Time, err error) {
	// Split hourMinute into slice with hours as 1st element and minutes as 2nd element.
	hourTimeParts := strings.Split(hourMinute, ":")
	now := time.Now().In(location)
	if len(hourTimeParts) < 2 {
		return dateTime, fmt.Errorf("got invalid hourMinute: %s", hourMinute)
	}
	hour, err := strconv.Atoi(hourTimeParts[0])
	if err != nil {
		return dateTime, fmt.Errorf("failed to parse hourMinute hour: %v, err: %w", hourTimeParts[0], err)
	}
	minute, err := strconv.Atoi(hourTimeParts[1])
	if err != nil {
		return dateTime, fmt.Errorf("failed to parse hourMinute minute: %v, err: %w", hourTimeParts[1], err)
	}
	dateTime = time.Date(
		now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, location,
	)

	// When triggerTime represents time in past correspondingly to time now
	// then adjust triggerTime to represents time for next day.
	if dateTime.Before(now) {
		dateTime = time.Date(
			now.Year(), now.Month(), now.Day()+1, hour, minute, 0, 0, location,
		)
	}
	return dateTime, nil
}

type groupedTriggersKey struct {
	RetailerID   string
	MfcID        string
	ScheduleType string
	TriggerAt    time.Time
}

// groupTriggersForPubSub groups TriggerItems with the same wpID, ScheduleType and TriggerAt,
// but different CutoffDateTime, to send one PubSub message for each group with multiple cutoffs
func (b *Business) groupForPubSub(triggerItems []r.TriggerItem) map[groupedTriggersKey][]r.TriggerItem {
	gropedTriggers := map[groupedTriggersKey][]r.TriggerItem{}

	for _, triggerItem := range triggerItems {
		key := groupedTriggersKey{
			RetailerID:   triggerItem.RetailerID,
			MfcID:        triggerItem.MfcID,
			ScheduleType: triggerItem.ScheduleType,
			TriggerAt:    triggerItem.TriggerAt,
		}
		if group, present := gropedTriggers[key]; present {
			group = append(group, triggerItem)
			gropedTriggers[key] = group
		} else {
			group = []r.TriggerItem{triggerItem}
			gropedTriggers[key] = group
		}
	}
	return gropedTriggers
}

func (b *Business) FireTriggers(ctx context.Context) (dto.FireTriggersResponse, error) {
	var firedAndPublishedTriggers []r.TriggerItem
	triggersToFire, err := b.repository.SelectTriggersToFire(ctx)
	if err != nil || len(triggersToFire) == 0 {
		return dto.FireTriggersResponse{Triggers: []dto.FireTriggerResponse{}}, err
	}
	groupedTriggers := b.groupForPubSub(triggersToFire)

	for key, group := range groupedTriggers {
		var cutoffs []time.Time
		for _, triggerItem := range group {
			cutoffs = append(cutoffs, triggerItem.CutoffDateTime)
		}

		pubsubAttributes := map[string]string{
			"env_type":    common.GetCtxEnv(ctx),
			"event_type":  "wp.TriggersFired",
			"retailer_id": key.RetailerID,
			"mfc_id":      key.MfcID,
			"source":      "wp",
		}

		message := dto.TriggersFiredEvent{
			ScheduleType: key.ScheduleType,
			Cutoffs:      cutoffs,
		}

		_, err := b.pubsub.PublishMessage(ctx, b.cfg.wpTopic, message, pubsubAttributes)
		if err != nil {
			b.logger.Errorw("Failed to publish fire trigger event",
				"retailer", key.RetailerID, "mfc", key.MfcID,
				"scheduleType", message.ScheduleType, "cutoffs", message.Cutoffs, zap.Error(err))
			return dto.FireTriggersResponse{}, err
		}

		b.logger.Infow("FireTrigger event successfully published ",
			"retailer", key.RetailerID, "mfc",
			key.MfcID, "schedule_type", message.ScheduleType, "cutoffs", message.Cutoffs)

		firedAt := time.Now()
		err = b.repository.MarkTriggersAsFired(ctx, group, firedAt)
		if err != nil {
			b.logger.Errorw("Failed to update triggers during fire",
				"retailer", key.RetailerID, "mfc", key.MfcID,
				"scheduleType", key.ScheduleType, "triggerAt", key.TriggerAt, zap.Error(err))
			return dto.FireTriggersResponse{}, err
		}
		for idx := range group {
			group[idx].FiredAt = &firedAt
		}
		firedAndPublishedTriggers = append(firedAndPublishedTriggers, group...)
	}
	return r.FiredTriggersFromSpannerToResponseView(firedAndPublishedTriggers), nil
}
