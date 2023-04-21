package business

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/to-com/wp/config"
	mockbusiness "github.com/to-com/wp/internal/business/mock"
	"github.com/to-com/wp/internal/business/validator"
	"github.com/to-com/wp/internal/common"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/mocks"
	"github.com/to-com/wp/internal/service/tsc"
	"github.com/to-com/wp/internal/testutils"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/to-com/wp/internal/repository"
)

func Test_generateTriggerForScheduleItem(ts *testing.T) {
	generationHoursLimit := 5

	timezone := "Asia/Dubai"
	location, err := time.LoadLocation(timezone)
	assert.Nil(ts, err)
	now := time.Now().Add(1 * time.Minute).In(location)
	hourMinute := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())

	defaultScheduleItem := repository.ScheduleItem{
		RetailerID:   "test_Retailer",
		MfcID:        "test_mfc",
		wpID:   "wave_plan_1",
		WaveID:       "wave_1",
		ScheduleID:   "schedule_1",
		Timezone:     timezone,
		ScheduleTime: hourMinute,
		CutoffTime:   hourMinute,
	}

	ts.Run("invalid Timezone", func(t *testing.T) {
		scheduleItem := defaultScheduleItem
		scheduleItem.Timezone = "error"

		_, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)

		assert.ErrorContains(t, err, "unknown time zone")
	})

	ts.Run("invalid ScheduleTime", func(t *testing.T) {
		scheduleItem := defaultScheduleItem
		scheduleItem.ScheduleTime = "invalid"

		_, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)

		assert.ErrorContains(t, err, "invalid hourMinute")
	})

	ts.Run("invalid CutoffTime", func(t *testing.T) {
		scheduleItem := defaultScheduleItem
		scheduleItem.CutoffTime = "broken"

		_, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)

		assert.ErrorContains(t, err, "failed to get cutoff date time")
	})

	ts.Run("simple success case", func(t *testing.T) {
		_, err := generateTriggerForScheduleItem(generationHoursLimit, defaultScheduleItem)
		assert.Nil(t, err)
	})
}

func Test_getTriggerTime(ts *testing.T) {
	timezone := "Europe/Kiev"
	location, err := time.LoadLocation(timezone)
	assert.Nil(ts, err)
	now := time.Now().In(location)

	ts.Run("invalid hourMinute string", func(t *testing.T) {
		_, err := getTriggerTime("error", location, 0)
		assert.ErrorContains(t, err, "got invalid hourMinute")
	})

	ts.Run("invalid hourMinute format", func(t *testing.T) {
		_, err := getTriggerTime("15H 27M", location, 0)
		assert.ErrorContains(t, err, "got invalid hourMinute")
	})

	ts.Run("invalid hourMinute value", func(t *testing.T) {
		_, err := getTriggerTime("HH:MM", location, 0)
		assert.ErrorContains(t, err, "failed to parse hourMinute")
	})

	ts.Run("invalid hourMinute minute", func(t *testing.T) {
		_, err := getTriggerTime("19:MM", location, 0)
		assert.ErrorContains(t, err, "failed to parse hourMinute")
	})

	ts.Run("exceeds upper bound", func(t *testing.T) {
		x := now.Add(-1 * time.Minute)
		hourMinute := fmt.Sprintf("%02d:%02d", x.Hour(), x.Minute())
		_, err := getTriggerTime(hourMinute, location, 0)

		assert.ErrorContains(t, err, "trigger time exceeds upper bound")
	})

	ts.Run("case with adjust trigger time", func(t *testing.T) {
		l, err := time.LoadLocation("America/Los_Angeles")
		assert.Nil(ts, err)
		now := time.Now().In(l)
		expected := now.Add(-1 * time.Minute)
		hourMinute := fmt.Sprintf("%02d:%02d", expected.Hour(), expected.Minute())

		actual, err := getTriggerTime(hourMinute, location, 48)

		assert.Nil(t, err)
		assert.Equal(t, expected.AddDate(0, 0, 1).Day(), actual.Day())
	})

	ts.Run("simple success case", func(t *testing.T) {
		hourMinute := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
		_, err = getTriggerTime(hourMinute, location, 48)

		assert.Nil(t, err)
	})
}

type wpSuite struct {
	suite.Suite

	ctrl           *gomock.Controller
	mockRepository *mockbusiness.MockDataStore
	mockPubsub     *mockbusiness.MockPublisher

	bs *Business

	tscURL string

	logs *observer.ObservedLogs
}

func TestBusinesswpSuite(t *testing.T) {
	suite.Run(t, new(wpSuite))
}

func (s *wpSuite) SetupTest() {
	cfg, err := config.Load()
	if err != nil {
		s.FailNow("unable to load config for tests, error: %v", err)
	}

	s.tscURL = cfg.TSCURLTemplate

	observedZapCore, logs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	s.logs = logs

	httpClient := &http.Client{}
	tscS := tsc.New(cfg, observedLogger, httpClient)

	s.ctrl = gomock.NewController(s.T())
	s.mockRepository = mockbusiness.NewMockDataStore(s.ctrl)
	s.mockPubsub = mockbusiness.NewMockPublisher(s.ctrl)

	s.bs = New(cfg, observedLogger, s.mockRepository, s.mockPubsub, tscS)
	s.bs.cfg.TriggersGenerationHoursLimit = 25
}

func (s *wpSuite) TearDownTest() {
	s.ctrl.Finish()
	gock.Off()
}

func (s *wpSuite) TestCreatewpWithNoSchedIspsOff() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpNoIspsRequest

	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)
	wpRequestSpannerViewAfter := toSpannerViewAfter(wpRequest)

	// Mock repository Createwp call and verify input params
	repCreatewp := s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Return(wpRequestSpannerViewAfter, nil).
		Times(1)

	// Generate expected pubsub data
	event, attrs := wpToPubsubInfo(ctx, repository.wpFromSpannerToResponseView(wpRequestSpannerViewAfter))

	// Mock pubsub PublishMessage call and verify input params
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqwpCreatedPubsubEvent(event), gomock.Eq(attrs)).
		Times(1).
		After(repCreatewp)

	// No CreateTriggers calls are expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Times(0)

	// Mock http call to TSC for ISPS_ENABLED value
	mocks.MockTscServiceIspsEnabledValue(s.tscURL, false)
	// Mock http call to TSC for location info
	mocks.MockTscLocationInfo(s.tscURL)

	// Test method call
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	testutils.Validatewp(s.T(), wpRequest, actualResponse)
}

func (s *wpSuite) TestCreatewpWithNoSchedIspsOn() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpNoIspsRequest

	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)
	wpRequestSpannerViewAfter := toSpannerViewAfter(wpRequest)

	// Mock repository Createwp call and verify input params
	repCreatewp := s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Return(wpRequestSpannerViewAfter, nil).
		Times(1)

	// Generate expected pubsub data
	event, attrs := wpToPubsubInfo(ctx, repository.wpFromSpannerToResponseView(wpRequestSpannerViewAfter))

	// Mock pubsub PublishMessage call and verify input params
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqwpCreatedPubsubEvent(event), gomock.Eq(attrs)).
		Times(1).
		After(repCreatewp)

	// No CreateTriggers calls are expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Times(0)

	// Mock http call to TSC for ISPS_ENABLED value
	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	// Mock http call to TSC for location info
	mocks.MockTscLocationInfo(s.tscURL)

	// Test method call
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	testutils.Validatewp(s.T(), wpRequest, actualResponse)

	allLogs := s.logs.All()
	s.Equal(1, len(allLogs))
	s.Equal(fmt.Sprintf("wp %v for Retailer: %v, Mfc: %v successfully published", event.wp.ID,
		event.wp.RetailerID, event.wp.MfcID), allLogs[0].Message)
}

func (s *wpSuite) TestCreatewpWithSchedIspsOn() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest

	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)
	wpRequestSpannerViewAfter := toSpannerViewAfter(wpRequest)

	// Mock repository Createwp call and verify input params
	repCreatewp := s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Return(wpRequestSpannerViewAfter, nil).
		Times(1)

	// Generate expected pubsub data
	event, attrs := wpToPubsubInfo(ctx, repository.wpFromSpannerToResponseView(wpRequestSpannerViewAfter))

	// Mock pubsub PublishMessage call and verify input params
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqwpCreatedPubsubEvent(event), gomock.Eq(attrs)).
		Times(1).
		After(repCreatewp)

	scheduleItemsFromwp := getwpScheduleItems(wpRequestSpannerViewAfter)
	triggers := make([]repository.Trigger, 0, len(scheduleItemsFromwp))
	for _, schedItem := range scheduleItemsFromwp {
		trigger, err := generateTriggerForScheduleItem(s.bs.cfg.TriggersGenerationHoursLimit, schedItem)
		if err != nil {
			s.T().Errorf("Fail to generate trigger for schedule item %v, error: %v", schedItem, err)
			s.T().FailNow()
		}
		triggers = append(triggers, trigger)
	}

	// Expected call to repository.CreateTriggers with generated triggers to make sure they are not changed by business part
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Eq(triggers)).
		Return(len(triggers), nil).
		Times(1)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	mocks.MockTscLocationInfo(s.tscURL)
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	testutils.Validatewp(s.T(), wpRequest, actualResponse)
}

func (s *wpSuite) TestCreatewpWithSchedIspsOff() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest

	// No repository Createwp call is expected
	s.mockRepository.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Times(0)

	// No pubsub PublishMessage call is expected
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// No repository.CreateTriggers call is expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Any(), gomock.Any()).
		Times(0)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, false)
	mocks.MockTscLocationInfo(s.tscURL)
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().NotEmpty(err)
	var validatorErr *validator.ValidationError
	s.IsType(validatorErr, err)
	s.Equal(dto.wpResponse{}, actualResponse)
}

func (s *wpSuite) TestCreatewpWithTscGetConfigsCallFail() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest

	// No repository Createwp call is expected
	s.mockRepository.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Times(0)

	// No pubsub PublishMessage call is expected
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// No repository.CreateTriggers call is expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Any(), gomock.Any()).
		Times(0)

	mocks.MockFailTscGetConfigItemsResponse(s.tscURL)
	mocks.MockTscLocationInfo(s.tscURL)
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().NotEmpty(err)
	s.Equal(dto.wpResponse{}, actualResponse)
	s.Regexp("an error occurred while retrieving service catalog config, response: .*", err.Error())
}

func (s *wpSuite) TestCreatewpWithTscGetLocationCallFail() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest

	// No repository Createwp call is expected
	s.mockRepository.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Times(0)

	// No pubsub PublishMessage call is expected
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// No repository.CreateTriggers call is expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Any(), gomock.Any()).
		Times(0)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	mocks.MockFailTscLocationInfo(s.tscURL)
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().NotEmpty(err)
	s.Regexp("failed to get timezone, err: .*", err.Error())
	s.Equal(dto.wpResponse{}, actualResponse)
}

func (s *wpSuite) TestCreatewpWithRepositoryCreatewpFail() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest
	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)

	// No repository Createwp call is expected
	s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Times(1).
		Return(repository.wp{}, fmt.Errorf("repository Createwp error"))

	// No pubsub PublishMessage call is expected
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// No repository.CreateTriggers call is expected
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Any(), gomock.Any()).
		Times(0)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	mocks.MockTscLocationInfo(s.tscURL)

	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().NotEmpty(err)
	s.Equal(dto.wpResponse{}, actualResponse)
	s.Equal(fmt.Errorf("repository Createwp error"), err)
}

func (s *wpSuite) TestCreatewpWithPublishwpFail() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest
	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)
	wpRequestSpannerViewAfter := toSpannerViewAfter(wpRequest)

	// Mock repository Createwp call and verify input params
	repCreatewp := s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Return(wpRequestSpannerViewAfter, nil).
		Times(1)

	// Generate expected pubsub data
	event, attrs := wpToPubsubInfo(ctx, repository.wpFromSpannerToResponseView(wpRequestSpannerViewAfter))

	// Mock pubsub PublishMessage call that returns error and verify input params
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqwpCreatedPubsubEvent(event), gomock.Eq(attrs)).
		Times(1).
		Return(nil, fmt.Errorf("fail to publish msg")).
		After(repCreatewp)

	scheduleItemsFromwp := getwpScheduleItems(wpRequestSpannerViewAfter)
	triggers := make([]repository.Trigger, 0, len(scheduleItemsFromwp))
	for _, schedItem := range scheduleItemsFromwp {
		trigger, err := generateTriggerForScheduleItem(s.bs.cfg.TriggersGenerationHoursLimit, schedItem)
		if err != nil {
			s.T().Errorf("Fail to generate trigger for schedule item %v, error: %v", schedItem, err)
			s.T().FailNow()
		}
		triggers = append(triggers, trigger)
	}

	// Expected call to repository.CreateTriggers with generated triggers to make sure they are not changed by business part
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Eq(triggers)).
		Return(len(triggers), nil).
		Times(1)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	mocks.MockTscLocationInfo(s.tscURL)

	s.T().Log("create wave plan")
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	testutils.Validatewp(s.T(), wpRequest, actualResponse)

	s.T().Log("validate logs")
	allLogs := s.logs.All()
	s.Equal(1, len(allLogs))
	s.Equal(fmt.Sprintf("Failed to publish wave plan: %s, error: fail to publish msg", event.wp.ID), allLogs[0].Message)
}

func (s *wpSuite) TestCreatewpWithRepositoryCreateTriggersFail() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	wpRequest := wpWithIspsRequest

	wpRequestSpannerViewBefore := toSpannerViewBefore(wpRequest)
	wpRequestSpannerViewAfter := toSpannerViewAfter(wpRequest)

	// Mock repository Createwp call and verify input params
	repCreatewp := s.mockRepository.
		EXPECT().
		Createwp(gomock.Eq(ctx), gomock.Eq(wpRequestSpannerViewBefore)).
		Return(wpRequestSpannerViewAfter, nil).
		Times(1)

	// Generate expected pubsub data
	event, attrs := wpToPubsubInfo(ctx, repository.wpFromSpannerToResponseView(wpRequestSpannerViewAfter))

	// Mock pubsub PublishMessage call and verify input params
	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqwpCreatedPubsubEvent(event), gomock.Eq(attrs)).
		Times(1).
		After(repCreatewp)

	// Expected call to repository.CreateTriggers with generated triggers to make sure they are not changed by business part
	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Return(0, fmt.Errorf("create triggers error")).
		Times(1)

	mocks.MockTscServiceIspsEnabledValue(s.tscURL, true)
	mocks.MockTscLocationInfo(s.tscURL)
	actualResponse, err := s.bs.Createwp(ctx, wpRequest)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	testutils.Validatewp(s.T(), wpRequest, actualResponse)

	s.T().Log("validate logs")
	allLogs := s.logs.All()
	s.Equal(3, len(allLogs))
	s.Equal(fmt.Sprintf("wp %v for Retailer: %v, Mfc: %v successfully published", event.wp.ID,
		event.wp.RetailerID, event.wp.MfcID), allLogs[0].Message)
	s.Equal("failed to save triggers to db: create triggers error", allLogs[1].Message)
	s.Equal("error occurred while creating triggers: create triggers error", allLogs[2].Message)
}

func (s *wpSuite) TestGetwpNoSchedules() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	// Mock repository Getwp call and verify input params
	s.mockRepository.
		EXPECT().
		Getwp(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return(wpNoIspsItems, nil).
		Times(1)

	// Test method call
	actualResponse, err := s.bs.Getwp(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	s.Equal(repository.wpItemsFromSpannerToResponseView(wpNoIspsItems), actualResponse)
}

func (s *wpSuite) TestGetwpWithSchedules() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	// Mock repository Getwp call and verify input params
	s.mockRepository.
		EXPECT().
		Getwp(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return(wpWithIspsItems, nil).
		Times(1)

	// Test method call
	actualResponse, err := s.bs.Getwp(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)

	s.T().Log("validate create wave plan response")
	s.Equal(repository.wpItemsFromSpannerToResponseView(wpWithIspsItems), actualResponse)
}

func (s *wpSuite) TestGetWaveEmptyPlan() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	// Mock repository Getwp call and verify input params
	s.mockRepository.
		EXPECT().
		Getwp(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return([]repository.wpItem{}, nil).
		Times(1)

	// Test method call
	actualResponse, err := s.bs.Getwp(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	s.Require().Empty(err)
	s.Require().Empty(actualResponse)
}

func (s *wpSuite) TestGetwpRepositoryReturnsError() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	// Mock repository Getwp call and verify input params
	s.mockRepository.
		EXPECT().
		Getwp(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return([]repository.wpItem{}, fmt.Errorf("get wave plan error")).
		Times(1)

	// Test method call
	actualResponse, err := s.bs.Getwp(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	s.Require().NotEmpty(err)
	s.Require().Empty(actualResponse)
	s.Equal(fmt.Errorf("get wave plan error"), err)
}

func (s *wpSuite) TestGenerateTriggers() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetScheduleItems(gomock.Eq(ctx)).
		Return(mockedSchedules, nil).
		Times(1)

	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Return(3, nil).
		Times(1)

	actualResponse, err := s.bs.GenerateTriggers(ctx)
	s.Require().Empty(err)
	s.Require().NotEmpty(actualResponse)
	s.Equal(dto.GenerateTriggersResponse{
		GeneratedTriggers: 3,
	}, actualResponse)
}

func (s *wpSuite) TestGenerateTriggersWithEmptySchedules() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetScheduleItems(gomock.Eq(ctx)).
		Return([]repository.ScheduleItem{}, nil).
		Times(1)

	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Times(0)

	actualResponse, err := s.bs.GenerateTriggers(ctx)
	s.Require().Empty(err)
	s.Equal(dto.GenerateTriggersResponse{
		GeneratedTriggers: 0,
	}, actualResponse)
}

func (s *wpSuite) TestGetScheduleItemsReturnsError() {
	errorMessage := "get schedule items error"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetScheduleItems(gomock.Eq(ctx)).
		Return([]repository.ScheduleItem{}, fmt.Errorf(errorMessage)).
		Times(1)

	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Times(0)

	actualResponse, err := s.bs.GenerateTriggers(ctx)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Require().Empty(actualResponse)
}

func (s *wpSuite) TestCreateTriggersReturnsError() {
	errorMessage := "create triggers error"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetScheduleItems(gomock.Eq(ctx)).
		Return(mockedSchedules, nil)

	s.mockRepository.
		EXPECT().
		CreateTriggers(gomock.Eq(ctx), gomock.Any()).
		Return(0, fmt.Errorf(errorMessage)).
		Times(1)

	actualResponse, err := s.bs.GenerateTriggers(ctx)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Require().Empty(actualResponse)
}

func (s *wpSuite) TestGetTriggers() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetTriggers(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return(triggerItems, nil).
		Times(1)

	actualResponse, err := s.bs.GetTriggers(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	expectedResponse := repository.TriggersFromSpannerToResponseView(triggerItems)
	s.Require().Empty(err)
	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestGetTriggersReturnsError() {
	errorMessage := "get triggers error"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		GetTriggers(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		Return([]repository.TriggerItem{}, fmt.Errorf(errorMessage)).
		Times(1)

	actualResponse, err := s.bs.GetTriggers(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Require().Empty(actualResponse)
}

func (s *wpSuite) TestGetTriggersEmpty() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")
	var emptyTriggers []repository.TriggerItem

	s.mockRepository.
		EXPECT().
		GetTriggers(gomock.Eq(ctx), gomock.Eq(mocks.GetMockLocationInfo().RetailerID), gomock.Eq(mocks.GetMockLocationInfo().MfcID)).
		Return(emptyTriggers, nil).
		Times(1)

	actualResponse, err := s.bs.GetTriggers(ctx, mocks.GetMockLocationInfo().RetailerID, mocks.GetMockLocationInfo().MfcID)
	expectedResponse := repository.TriggersFromSpannerToResponseView(emptyTriggers)
	s.Require().Empty(err)
	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestFireTriggers() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	repSelectTriggersToFire := s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return(triggerItems2, nil).
		Times(1)

	for key, val := range groupedTriggers {
		event, attrs := fireTriggersToPubsubInfo(ctx, key, val)
		s.mockPubsub.
			EXPECT().
			PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqPubsubFireTriggerEvent(event), gomock.Eq(attrs)).
			Times(1).
			After(repSelectTriggersToFire)

		s.mockRepository.
			EXPECT().
			MarkTriggersAsFired(gomock.Eq(ctx), gomock.Eq(val), gomock.Any()).
			Return(nil).
			Times(1)
	}

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Empty(err)
	s.NotEmpty(actualResponse.Triggers[0].FiredAt)
	var triggers []dto.FireTriggerResponse
	for _, tr := range actualResponse.Triggers {
		tr.FiredAt = nil
		triggers = append(triggers, tr)
	}

	expectedResponse := repository.FiredTriggersFromSpannerToResponseView(triggerItems2)
	s.ElementsMatch(expectedResponse.Triggers, triggers)
}

func (s *wpSuite) TestSelectTriggersToFireReturnsError() {
	errorMessage := "select trigger to fire error"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return(make([]repository.TriggerItem, 0), fmt.Errorf(errorMessage)).
		Times(1)

	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Equal(dto.FireTriggersResponse{Triggers: make([]dto.FireTriggerResponse, 0)}, actualResponse)
}

func (s *wpSuite) TestPublishToPubSubReturnsError() {
	errorMessage := "failed publishing to pubsub"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return(triggerItems2, nil).
		Times(1)

	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), gomock.Any(), gomock.Any()).
		Return(nil, fmt.Errorf(errorMessage)).
		Times(1)

	s.mockRepository.
		EXPECT().
		MarkTriggersAsFired(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		Return(nil).
		Times(0)

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Equal(dto.FireTriggersResponse{}, actualResponse)
}

func (s *wpSuite) TestMarkTriggersAsFiredReturnsError() {
	errorMessage := "failed firing triggers"
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return(triggerItems2, nil).
		Times(1)

	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), gomock.Any(), gomock.Any()).
		Times(1)

	s.mockRepository.
		EXPECT().
		MarkTriggersAsFired(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		Return(fmt.Errorf(errorMessage)).
		Times(1)

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Require().NotEmpty(err)
	s.Equal(errorMessage, err.Error())
	s.Equal(dto.FireTriggersResponse{}, actualResponse)
}

func (s *wpSuite) TestNoTriggersToFire() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return([]repository.TriggerItem{}, nil).
		Times(1)

	s.mockPubsub.
		EXPECT().
		PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), gomock.Any(), gomock.Any()).
		Times(0)

	s.mockRepository.
		EXPECT().
		MarkTriggersAsFired(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		Return(nil).
		Times(0)

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Empty(err)
	s.Equal(dto.FireTriggersResponse{Triggers: []dto.FireTriggerResponse{}}, actualResponse)
}

func (s *wpSuite) TestFireTriggersForDifferentRetailers() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")

	repSelectTriggersToFire := s.mockRepository.
		EXPECT().
		SelectTriggersToFire(gomock.Eq(ctx)).
		Return(triggerItems3, nil).
		Times(1)

	for key, val := range groupedTriggers2 {
		event, attrs := fireTriggersToPubsubInfo(ctx, key, val)
		s.mockPubsub.
			EXPECT().
			PublishMessage(gomock.Eq(ctx), gomock.Eq(s.bs.cfg.wpTopic), EqPubsubFireTriggerEvent(event), gomock.Eq(attrs)).
			Times(1).
			After(repSelectTriggersToFire)

		s.mockRepository.
			EXPECT().
			MarkTriggersAsFired(gomock.Eq(ctx), gomock.Eq(val), gomock.Any()).
			Return(nil).
			Times(1)
	}

	actualResponse, err := s.bs.FireTriggers(ctx)
	s.Empty(err)
	s.NotEmpty(actualResponse.Triggers[0].FiredAt)
	var triggers []dto.FireTriggerResponse
	for _, tr := range actualResponse.Triggers {
		tr.FiredAt = nil
		triggers = append(triggers, tr)
	}

	expectedResponse := repository.FiredTriggersFromSpannerToResponseView(triggerItems3)
	s.ElementsMatch(expectedResponse.Triggers, triggers)
}

func (s *wpSuite) TestGroupForPubSub() {
	s.Run("group triggers by trigger time", func() {
		grouped := s.bs.groupForPubSub(triggerItems2)
		s.Equal(grouped, groupedTriggers)
	})

	s.Run("group triggers fof different retailers", func() {
		grouped := s.bs.groupForPubSub(triggerItems3)
		s.Equal(grouped, groupedTriggers2)
	})
}

func (s *wpSuite) TestGenerateTriggersForScheduleItems() {
	actualTriggers, err := s.bs.generateTriggersForScheduleItems(mockedSchedules)
	s.Empty(err)
	s.ElementsMatch(actualTriggers, scheduleItemsToTriggers(mockedSchedules))
}

func (s *wpSuite) TestGenerateTriggerForScheduleItem() {
	generationHoursLimit := 24

	timezone := "Asia/Dubai"
	location, err := time.LoadLocation(timezone)
	s.Nil(err)
	now := time.Now().Add(1 * time.Minute).In(location)

	hourMinute := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())

	defaultScheduleItem := repository.ScheduleItem{
		RetailerID:   "test_Retailer",
		MfcID:        "test_mfc",
		wpID:   "wave_plan_1",
		WaveID:       "wave_1",
		ScheduleID:   "schedule_1",
		Timezone:     timezone,
		ScheduleTime: hourMinute,
		CutoffTime:   hourMinute,
	}

	s.Run("schedule time is later today, cutoff is tomorrow > 24h from now", func() {
		scheduleItem := defaultScheduleItem
		scheduleTime := now.Add(3 * time.Minute)
		cutoffTime := now.Add(1 * time.Minute)
		scheduleItem.ScheduleTime = fmt.Sprintf("%02d:%02d", scheduleTime.Hour(), scheduleTime.Minute())
		scheduleItem.CutoffTime = fmt.Sprintf("%02d:%02d", cutoffTime.Hour(), cutoffTime.Minute())
		expectedCutoffDateTime := cutoffTime.Add(24 * time.Hour).UTC().Format("2006-01-02 15:04")
		expectedTriggerAt := scheduleTime.UTC().Format("2006-01-02 15:04")

		trigger, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)
		s.Nil(err)
		s.Equal(expectedTriggerAt, trigger.TriggerAt.Format("2006-01-02 15:04"))
		s.Equal(expectedCutoffDateTime, trigger.CutoffDateTime.Format("2006-01-02 15:04"))
	})

	s.Run("schedule time is later today, cutoff is tomorrow < 24h from now", func() {
		scheduleItem := defaultScheduleItem
		scheduleTime := now.Add(1 * time.Minute)
		cutoffTime := now.Add(-1 * time.Minute)
		scheduleItem.ScheduleTime = fmt.Sprintf("%02d:%02d", scheduleTime.Hour(), scheduleTime.Minute())
		scheduleItem.CutoffTime = fmt.Sprintf("%02d:%02d", cutoffTime.Hour(), cutoffTime.Minute())
		expectedCutoffDateTime := cutoffTime.Add(24 * time.Hour).UTC().Format("2006-01-02 15:04")
		expectedTriggerAt := scheduleTime.UTC().Format("2006-01-02 15:04")

		trigger, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)
		s.Nil(err)
		s.Equal(expectedTriggerAt, trigger.TriggerAt.Format("2006-01-02 15:04"))
		s.Equal(expectedCutoffDateTime, trigger.CutoffDateTime.Format("2006-01-02 15:04"))
	})

	s.Run("schedule time tomorrow, cutoff is the next day", func() {
		scheduleItem := defaultScheduleItem
		scheduleTime := now.Add(-1 * time.Minute)
		cutoffTime := now.Add(-2 * time.Minute)
		scheduleItem.ScheduleTime = fmt.Sprintf("%02d:%02d", scheduleTime.Hour(), scheduleTime.Minute())
		scheduleItem.CutoffTime = fmt.Sprintf("%02d:%02d", cutoffTime.Hour(), cutoffTime.Minute())
		expectedCutoffDateTime := cutoffTime.Add(48 * time.Hour).UTC().Format("2006-01-02 15:04")
		expectedTriggerAt := scheduleTime.Add(24 * time.Hour).UTC().Format("2006-01-02 15:04")

		trigger, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)
		s.Nil(err)
		s.Equal(expectedTriggerAt, trigger.TriggerAt.Format("2006-01-02 15:04"))
		s.Equal(expectedCutoffDateTime, trigger.CutoffDateTime.Format("2006-01-02 15:04"))
	})

	s.Run("cutoff and schedule time are tomorrow", func() {
		scheduleItem := defaultScheduleItem
		scheduleTime := now.Add(-1 * time.Minute)
		cutoffTime := now.Add(1 * time.Minute)
		scheduleItem.ScheduleTime = fmt.Sprintf("%02d:%02d", scheduleTime.Hour(), scheduleTime.Minute())
		scheduleItem.CutoffTime = fmt.Sprintf("%02d:%02d", cutoffTime.Hour(), cutoffTime.Minute())
		expectedCutoffDateTime := cutoffTime.Add(24 * time.Hour).UTC().Format("2006-01-02 15:04")
		expectedTriggerAt := scheduleTime.Add(24 * time.Hour).UTC().Format("2006-01-02 15:04")

		trigger, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)
		s.Nil(err)
		s.Equal(expectedTriggerAt, trigger.TriggerAt.Format("2006-01-02 15:04"))
		s.Equal(expectedCutoffDateTime, trigger.CutoffDateTime.Format("2006-01-02 15:04"))
	})

	s.Run("cutoff and schedule time are later today", func() {
		scheduleItem := defaultScheduleItem
		scheduleTime := now.Add(1 * time.Minute)
		cutoffTime := now.Add(3 * time.Minute)
		scheduleItem.ScheduleTime = fmt.Sprintf("%02d:%02d", scheduleTime.Hour(), scheduleTime.Minute())
		scheduleItem.CutoffTime = fmt.Sprintf("%02d:%02d", cutoffTime.Hour(), cutoffTime.Minute())
		expectedCutoffDateTime := cutoffTime.UTC().Format("2006-01-02 15:04")
		expectedTriggerAt := scheduleTime.UTC().Format("2006-01-02 15:04")

		trigger, err := generateTriggerForScheduleItem(generationHoursLimit, scheduleItem)
		s.Nil(err)
		s.Equal(expectedTriggerAt, trigger.TriggerAt.Format("2006-01-02 15:04"))
		s.Equal(expectedCutoffDateTime, trigger.CutoffDateTime.Format("2006-01-02 15:04"))
	})
}
