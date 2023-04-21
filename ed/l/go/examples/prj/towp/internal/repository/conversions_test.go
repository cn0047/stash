package repository

import (
	"cloud.google.com/go/spanner"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/internal/dto"
	"testing"
	"time"
)

func mockNewSpannerID() string {
	return "9f34942f-6509-4b4c-a286-06c03887c050"
}

func mockNowUTC() time.Time {
	return time.Date(2022, time.October, 12, 10, 25, 37, 378112000, time.UTC)
}

func spannerString(v string) spanner.NullString {
	return spanner.NullString{StringVal: v, Valid: true}
}

var retailer, mfc, user = "FAKE-RETAILER", "fake-mfc", "JohnDoe"

var wpSpanner = wp{
	ID:         mockNewSpannerID(),
	RetailerID: retailer,
	MfcID:      mfc,
	Timezone:   "Asia/Dubai",
	CreatedAt:  mockNowUTC(),
	CreatedBy:  user,
	Waves: []Wave{
		{
			wpID: mockNewSpannerID(),
			ID:         mockNewSpannerID(),
			Cutoff:     "09:00",
			FromTime:   "00:00",
			ToTime:     "16:00",
			Schedules: []Schedule{
				{
					ID:           mockNewSpannerID(),
					ScheduleType: "prelim_picklist",
					ScheduleTime: "12:00",
				},
				{
					ID:           mockNewSpannerID(),
					ScheduleType: "delta_picklist",
					ScheduleTime: "13:00",
				},
			},
		},
	},
}

var wpResponse = dto.wpResponse{
	ID:         mockNewSpannerID(),
	RetailerID: retailer,
	MfcID:      mfc,
	Timezone:   "Asia/Dubai",
	CreatedAt:  mockNowUTC(),
	CreatedBy:  user,
	Waves: []dto.WaveResponse{
		{
			ID:       mockNewSpannerID(),
			Cutoff:   "09:00",
			FromTime: "00:00",
			ToTime:   "16:00",
			Schedules: []dto.ScheduleResponse{
				{
					ID:           mockNewSpannerID(),
					ScheduleType: "prelim_picklist",
					ScheduleTime: "12:00",
				},
				{
					ID:           mockNewSpannerID(),
					ScheduleType: "delta_picklist",
					ScheduleTime: "13:00",
				},
			},
		},
	},
}

func TestwpFromSpannerToResponseView(t *testing.T) {
	got := wpFromSpannerToResponseView(wpSpanner)
	assert.Equal(t, wpResponse, got)
}

func TestwpItemsFromSpannerToResponseView(t *testing.T) {
	wpItems := []wpItem{
		{
			wpID:   mockNewSpannerID(),
			RetailerID:   retailer,
			MfcID:        mfc,
			Timezone:     "Asia/Dubai",
			CreatedAt:    mockNowUTC(),
			CreatedBy:    user,
			WaveID:       mockNewSpannerID(),
			Cutoff:       "09:00",
			FromTime:     "00:00",
			ToTime:       "16:00",
			ScheduleID:   spannerString(mockNewSpannerID()),
			ScheduleType: spannerString("prelim_picklist"),
			ScheduleTime: spannerString("12:00"),
		},
		{
			wpID:   mockNewSpannerID(),
			RetailerID:   retailer,
			MfcID:        mfc,
			Timezone:     "Asia/Dubai",
			CreatedAt:    mockNowUTC(),
			CreatedBy:    user,
			WaveID:       mockNewSpannerID(),
			Cutoff:       "09:00",
			FromTime:     "00:00",
			ToTime:       "16:00",
			ScheduleID:   spannerString(mockNewSpannerID()),
			ScheduleType: spannerString("delta_picklist"),
			ScheduleTime: spannerString("13:00"),
		},
	}

	got := wpItemsFromSpannerToResponseView(wpItems)

	assert.Equal(t, wpResponse, got)
}

func TestwpToSpannerView(t *testing.T) {
	wpReq := dto.wpRequest{
		RetailerID: retailer,
		MfcID:      mfc,
		UserID:     user,
		Waves: []dto.WaveRequest{
			{
				Cutoff:     "09:00",
				FromTime:   "00:00",
				ToTime:     "16:00",
				PrelimTime: "12:00",
				DeltaTime:  "13:00",
			},
		},
	}

	got := wpToSpannerView(wpReq)

	assert.Equal(t, "fake-retailer", got.RetailerID)
	assert.Equal(t, wpSpanner.MfcID, got.MfcID)
	assert.Equal(t, wpSpanner.CreatedBy, got.CreatedBy)
}

func TestTriggersFromSpannerToResponseView(t *testing.T) {
	now := mockNowUTC()
	spannerTriggers := []TriggerItem{
		{
			RetailerID:     retailer,
			MfcID:          mfc,
			Timezone:       "Europe/Kiev",
			wpID:     mockNewSpannerID(),
			WaveID:         mockNewSpannerID(),
			Cutoff:         "14:00",
			ScheduleID:     mockNewSpannerID(),
			ScheduleTime:   "13:00",
			TriggerAt:      mockNowUTC(),
			CreatedAt:      mockNowUTC(),
			CutoffDateTime: now,
			FiredAt:        &now,
		},
	}

	got := TriggersFromSpannerToResponseView(spannerTriggers)

	want := dto.GetTriggersResponse{
		Triggers: []dto.TriggerResponse{
			{
				RetailerID:     retailer,
				MfcID:          mfc,
				Timezone:       "Europe/Kiev",
				wpID:     mockNewSpannerID(),
				WaveID:         mockNewSpannerID(),
				Cutoff:         "14:00",
				ScheduleID:     mockNewSpannerID(),
				ScheduleTime:   "13:00",
				TriggerAt:      mockNowUTC(),
				CreatedAt:      mockNowUTC(),
				CutoffDatetime: now,
				FiredAt:        &now,
			},
		},
	}

	assert.Equal(t, want, got)
}
