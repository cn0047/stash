package business

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/to-com/wp/internal/common"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/mocks"
	"github.com/to-com/wp/internal/repository"
	"time"
)

var (
	fakeMfc               = "fake-mfc"
	fakeRetailer          = "fake-retailer"
	wpNoIspsRequest = dto.wpRequest{
		RetailerID: mocks.GetMockLocationInfo().RetailerID,
		MfcID:      mocks.GetMockLocationInfo().MfcID,
		UserID:     mocks.GetUserPermissionMock()["user_id"],
		Waves: []dto.WaveRequest{
			{
				Cutoff:   "08:00",
				FromTime: "09:00",
				ToTime:   "11:59",
			},
			{
				Cutoff:   "10:00",
				FromTime: "12:00",
				ToTime:   "16:59",
			},
			{
				Cutoff:   "15:00",
				FromTime: "17:00",
				ToTime:   "08:59",
			},
		},
	}

	wpWithIspsRequest = dto.wpRequest{
		RetailerID: mocks.GetMockLocationInfo().RetailerID,
		MfcID:      mocks.GetMockLocationInfo().MfcID,
		UserID:     mocks.GetUserPermissionMock()["user_id"],
		Waves: []dto.WaveRequest{
			{
				Cutoff:     "08:00",
				FromTime:   "09:00",
				ToTime:     "11:59",
				PrelimTime: "04:00",
				DeltaTime:  "07:00",
			},
			{
				Cutoff:     "10:00",
				FromTime:   "12:00",
				ToTime:     "16:59",
				PrelimTime: "04:00",
			},
			{
				Cutoff:    "15:00",
				FromTime:  "17:00",
				ToTime:    "08:59",
				DeltaTime: "14:00",
			},
		},
	}

	wpID          = uuid.NewString()
	wpNoIspsItems = []repository.wpItem{
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "08:00",
			FromTime:   "09:00",
			ToTime:     "11:59",
		},
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "10:00",
			FromTime:   "12:00",
			ToTime:     "16:59",
		},
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "15:00",
			FromTime:   "17:00",
			ToTime:     "08:59",
		},
	}

	wpWithIspsItems = []repository.wpItem{
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "08:00",
			FromTime:   "09:00",
			ToTime:     "11:59",
			ScheduleID: spanner.NullString{
				StringVal: uuid.NewString(),
				Valid:     true,
			},
			ScheduleType: spanner.NullString{
				StringVal: dto.Prelim,
				Valid:     true,
			},
			ScheduleTime: spanner.NullString{
				StringVal: "04:00",
				Valid:     true,
			},
		},
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "08:00",
			FromTime:   "09:00",
			ToTime:     "11:59",
			ScheduleID: spanner.NullString{
				StringVal: uuid.NewString(),
				Valid:     true,
			},
			ScheduleType: spanner.NullString{
				StringVal: dto.Delta,
				Valid:     true,
			},
			ScheduleTime: spanner.NullString{
				StringVal: "07:00",
				Valid:     true,
			},
		},
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "10:00",
			FromTime:   "12:00",
			ToTime:     "16:59",
			ScheduleID: spanner.NullString{
				StringVal: uuid.NewString(),
				Valid:     true,
			},
			ScheduleType: spanner.NullString{
				StringVal: dto.Prelim,
				Valid:     true,
			},
			ScheduleTime: spanner.NullString{
				StringVal: "04:00",
				Valid:     true,
			},
		},
		{
			wpID: wpID,
			RetailerID: mocks.GetMockLocationInfo().RetailerID,
			MfcID:      mocks.GetMockLocationInfo().MfcID,
			Timezone:   mocks.GetMockLocationInfo().Timezone,
			CreatedAt:  time.Now(),
			CreatedBy:  mocks.GetUserPermissionMock()["user_id"],
			WaveID:     uuid.NewString(),
			Cutoff:     "15:00",
			FromTime:   "17:00",
			ToTime:     "08:59",
			ScheduleID: spanner.NullString{
				StringVal: uuid.NewString(),
				Valid:     true,
			},
			ScheduleType: spanner.NullString{
				StringVal: dto.Delta,
				Valid:     true,
			},
			ScheduleTime: spanner.NullString{
				StringVal: "14:00",
				Valid:     true,
			},
		},
	}

	loc, _       = time.LoadLocation(mocks.GetMockLocationInfo().Timezone)
	triggerTime1 = time.Now().In(loc).Add(1 * time.Hour)
	cutoff1      = triggerTime1.Add(4 * time.Hour)
	triggerTime2 = time.Now().In(loc).Add(2 * time.Hour)
	cutoff2      = triggerTime2.Add(4 * time.Hour)

	mockedSchedules = []repository.ScheduleItem{
		{
			RetailerID:   mocks.GetMockLocationInfo().RetailerID,
			MfcID:        mocks.GetMockLocationInfo().MfcID,
			Timezone:     mocks.GetMockLocationInfo().Timezone,
			wpID:   uuid.NewString(),
			WaveID:       uuid.NewString(),
			ScheduleID:   uuid.NewString(),
			ScheduleTime: fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			CutoffTime:   fmt.Sprintf("%v:%v", cutoff1.Hour(), cutoff1.Minute()),
			TriggerAt:    &triggerTime1,
		},
		{
			RetailerID:   mocks.GetMockLocationInfo().RetailerID,
			MfcID:        mocks.GetMockLocationInfo().MfcID,
			Timezone:     mocks.GetMockLocationInfo().Timezone,
			wpID:   uuid.NewString(),
			WaveID:       uuid.NewString(),
			ScheduleID:   uuid.NewString(),
			ScheduleTime: fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			CutoffTime:   fmt.Sprintf("%v:%v", cutoff2.Hour(), cutoff2.Minute()),
			TriggerAt:    &triggerTime1,
		},
		{
			RetailerID:   mocks.GetMockLocationInfo().RetailerID + "1",
			MfcID:        mocks.GetMockLocationInfo().MfcID,
			Timezone:     mocks.GetMockLocationInfo().Timezone,
			wpID:   uuid.NewString(),
			WaveID:       uuid.NewString(),
			ScheduleID:   uuid.NewString(),
			ScheduleTime: fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			CutoffTime:   fmt.Sprintf("%v:%v", cutoff1.Hour(), cutoff1.Minute()),
			TriggerAt:    &triggerTime2,
		},
		{
			RetailerID:   mocks.GetMockLocationInfo().RetailerID,
			MfcID:        mocks.GetMockLocationInfo().MfcID + "1",
			Timezone:     mocks.GetMockLocationInfo().Timezone,
			wpID:   uuid.NewString(),
			WaveID:       uuid.NewString(),
			ScheduleID:   uuid.NewString(),
			ScheduleTime: fmt.Sprintf("%v:%v", triggerTime2.Hour(), triggerTime2.Minute()),
			CutoffTime:   fmt.Sprintf("%v:%v", cutoff2.Hour(), cutoff2.Minute()),
			TriggerAt:    &triggerTime2,
		},
	}
	triggerItems = []repository.TriggerItem{
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "11:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Prelim",
			TriggerAt:      time.Now(),
			CreatedAt:      time.Now(),
			CutoffDateTime: time.Now().Add(+5 * time.Hour),
			FiredAt:        nil,
		},
	}
	timeNow       = time.Now()
	timeNowPlus5  = timeNow.Add(+5 * time.Hour)
	timeNowPlus4  = timeNow.Add(+4 * time.Hour)
	timeNowPlus6  = timeNow.Add(+6 * time.Hour)
	triggerItems2 = []repository.TriggerItem{
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "11:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Prelim",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus5,
			FiredAt:        nil,
		},
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "10:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Prelim",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus4,
			FiredAt:        nil,
		},
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "11:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Delta",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus6,
			FiredAt:        nil,
		},
	}
	key1            = groupedTriggersKey{RetailerID: fakeRetailer, MfcID: fakeMfc, ScheduleType: "Prelim", TriggerAt: triggerTime1}
	key2            = groupedTriggersKey{RetailerID: fakeRetailer, MfcID: fakeMfc, ScheduleType: "Delta", TriggerAt: triggerTime1}
	groupedTriggers = map[groupedTriggersKey][]repository.TriggerItem{
		key1: {
			triggerItems2[0],
			triggerItems2[1]},
		key2: {triggerItems2[2]},
	}
	triggerItems3 = []repository.TriggerItem{
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID + "1",
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "11:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Prelim",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus5,
			FiredAt:        nil,
		},
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID + "1",
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "10:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Prelim",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus4,
			FiredAt:        nil,
		},
		{
			RetailerID:     mocks.GetMockLocationInfo().RetailerID,
			MfcID:          mocks.GetMockLocationInfo().MfcID,
			Timezone:       mocks.GetMockLocationInfo().Timezone,
			wpID:     uuid.NewString(),
			WaveID:         uuid.NewString(),
			Cutoff:         "11:00",
			ScheduleID:     uuid.NewString(),
			ScheduleTime:   fmt.Sprintf("%v:%v", triggerTime1.Hour(), triggerTime1.Minute()),
			ScheduleType:   "Delta",
			TriggerAt:      triggerTime1,
			CreatedAt:      timeNow,
			CutoffDateTime: timeNowPlus6,
			FiredAt:        nil,
		},
	}
	key3             = groupedTriggersKey{RetailerID: fakeRetailer, MfcID: fakeMfc + "1", ScheduleType: "Prelim", TriggerAt: triggerTime1}
	key4             = groupedTriggersKey{RetailerID: fakeRetailer + "1", MfcID: fakeMfc, ScheduleType: "Prelim", TriggerAt: triggerTime1}
	key5             = groupedTriggersKey{RetailerID: fakeRetailer, MfcID: fakeMfc, ScheduleType: "Delta", TriggerAt: triggerTime1}
	groupedTriggers2 = map[groupedTriggersKey][]repository.TriggerItem{
		key3: {triggerItems3[0]},
		key4: {triggerItems3[1]},
		key5: {triggerItems3[2]},
	}
)

func toSpannerViewBefore(wpRequest dto.wpRequest) repository.wp {
	wpRequestSpannerViewBefore := repository.wpToSpannerView(wpRequest)
	wpRequestSpannerViewBefore.Timezone = mocks.GetMockLocationInfo().Timezone
	return wpRequestSpannerViewBefore
}

func toSpannerViewAfter(wpRequest dto.wpRequest) repository.wp {
	wpRequestSpannerViewAfter := repository.wpToSpannerView(wpRequest)
	wpRequestSpannerViewAfter.ID = uuid.NewString()
	wpRequestSpannerViewAfter.CreatedAt = time.Now()
	wpRequestSpannerViewAfter.Timezone = mocks.GetMockLocationInfo().Timezone
	for i := 0; i < len(wpRequestSpannerViewAfter.Waves); i++ {
		wpRequestSpannerViewAfter.Waves[i].ID = uuid.NewString()
		wpRequestSpannerViewAfter.Waves[i].wpID = wpRequestSpannerViewAfter.ID
		for j := 0; j < len(wpRequestSpannerViewAfter.Waves[i].Schedules); j++ {
			wpRequestSpannerViewAfter.Waves[i].Schedules[j].ID = uuid.NewString()
			wpRequestSpannerViewAfter.Waves[i].Schedules[j].WaveID = wpRequestSpannerViewAfter.Waves[i].ID
			wpRequestSpannerViewAfter.Waves[i].Schedules[j].wpID = wpRequestSpannerViewAfter.Waves[i].wpID
		}
	}
	return wpRequestSpannerViewAfter
}

func wpToPubsubInfo(ctx context.Context, resp dto.wpResponse) (dto.wpCreatedEvent, map[string]string) {
	pubsubAttributes := map[string]string{
		"env_type":    common.GetCtxEnv(ctx),
		"event_type":  "wp.Created",
		"retailer_id": resp.RetailerID,
		"mfc_id":      resp.MfcID,
		"source":      "wp",
	}

	message := dto.wpCreatedEvent{
		CreatedAt: time.Now(),
		wp:  resp,
	}

	return message, pubsubAttributes
}

func fireTriggersToPubsubInfo(ctx context.Context,
	triggerKey groupedTriggersKey,
	triggerItems []repository.TriggerItem) (dto.TriggersFiredEvent, map[string]string) {
	var cutoffs []time.Time
	for _, triggerItem := range triggerItems {
		cutoffs = append(cutoffs, triggerItem.CutoffDateTime)
	}

	pubsubAttributes := map[string]string{
		"env_type":    common.GetCtxEnv(ctx),
		"event_type":  "wp.TriggersFired",
		"retailer_id": triggerKey.RetailerID,
		"mfc_id":      triggerKey.MfcID,
		"source":      "wp",
	}

	message := dto.TriggersFiredEvent{
		ScheduleType: triggerKey.ScheduleType,
		Cutoffs:      cutoffs,
	}
	return message, pubsubAttributes
}

func getwpScheduleItems(wp repository.wp) []repository.ScheduleItem {
	scheduleItems := make([]repository.ScheduleItem, 0)
	for _, wave := range wp.Waves {
		for _, schedule := range wave.Schedules {
			scheduleItem := repository.ScheduleItem{
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

	return scheduleItems
}

func scheduleItemsToTriggers(scheduleItems []repository.ScheduleItem) (triggers []repository.Trigger) {
	for _, scheduleItem := range scheduleItems {
		location, _ := time.LoadLocation(scheduleItem.Timezone)
		triggerAt, _ := convertTimeToDateTime(scheduleItem.ScheduleTime, location)
		cutoffDateTime, _ := convertTimeToDateTime(scheduleItem.CutoffTime, location)
		trigger := repository.Trigger{
			wpID:     scheduleItem.wpID,
			WaveID:         scheduleItem.WaveID,
			ScheduleID:     scheduleItem.ScheduleID,
			TriggerAt:      triggerAt.UTC(), // IMPORTANT: time in UTC.
			CutoffDateTime: cutoffDateTime.UTC(),
		}
		triggers = append(triggers, trigger)
	}
	return triggers
}
