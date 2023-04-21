package repository

import (
	"github.com/to-com/wp/internal/dto"
	"strings"
)

func wpFromSpannerToResponseView(wp wp) dto.wpResponse {
	wpResponse := dto.wpResponse{
		ID:         wp.ID,
		RetailerID: wp.RetailerID,
		MfcID:      wp.MfcID,
		Timezone:   wp.Timezone,
		CreatedAt:  wp.CreatedAt,
		CreatedBy:  wp.CreatedBy,
		Waves:      make([]dto.WaveResponse, 0, len(wp.Waves)),
	}

	for _, wave := range wp.Waves {
		waveResponse := dto.WaveResponse{
			ID:        wave.ID,
			Cutoff:    wave.Cutoff,
			FromTime:  wave.FromTime,
			ToTime:    wave.ToTime,
			Schedules: make([]dto.ScheduleResponse, 0, len(wave.Schedules)),
		}

		for _, schedule := range wave.Schedules {
			scheduleResponse := dto.ScheduleResponse{
				ID:           schedule.ID,
				ScheduleType: schedule.ScheduleType,
				ScheduleTime: schedule.ScheduleTime,
			}
			waveResponse.Schedules = append(waveResponse.Schedules, scheduleResponse)
		}

		wpResponse.Waves = append(wpResponse.Waves, waveResponse)
	}

	return wpResponse
}

func wpItemsFromSpannerToResponseView(wpItems []wpItem) dto.wpResponse {
	wpResponse := dto.wpResponse{
		ID:         wpItems[0].wpID,
		RetailerID: wpItems[0].RetailerID,
		MfcID:      wpItems[0].MfcID,
		Timezone:   wpItems[0].Timezone,
		CreatedAt:  wpItems[0].CreatedAt,
		CreatedBy:  wpItems[0].CreatedBy,
		Waves:      make([]dto.WaveResponse, 0),
	}

	waves := make(map[string]*dto.WaveResponse, 0)

	for _, item := range wpItems {
		if _, ok := waves[item.WaveID]; !ok {
			wave := dto.WaveResponse{
				ID:        item.WaveID,
				Cutoff:    item.Cutoff,
				FromTime:  item.FromTime,
				ToTime:    item.ToTime,
				Schedules: make([]dto.ScheduleResponse, 0),
			}
			wpResponse.Waves = append(wpResponse.Waves, wave)
			waves[item.WaveID] = &wpResponse.Waves[len(wpResponse.Waves)-1]
		}

		if item.ScheduleID.Valid {
			waves[item.WaveID].Schedules = append(waves[item.WaveID].Schedules, dto.ScheduleResponse{
				ID:           item.ScheduleID.StringVal,
				ScheduleType: item.ScheduleType.StringVal,
				ScheduleTime: item.ScheduleTime.StringVal,
			})
		}
	}

	return wpResponse
}

func wpToSpannerView(wp dto.wpRequest) wp {
	wpDB := wp{
		RetailerID: strings.ToLower(wp.RetailerID),
		MfcID:      wp.MfcID,
		CreatedBy:  wp.UserID,
		Waves:      make([]Wave, 0, len(wp.Waves)),
	}

	for _, wave := range wp.Waves {
		waveDB := Wave{
			Cutoff:    wave.Cutoff,
			FromTime:  wave.FromTime,
			ToTime:    wave.ToTime,
			Schedules: make([]Schedule, 0),
		}
		if wave.PrelimTime != "" {
			schedule := Schedule{
				ScheduleTime: wave.PrelimTime,
				ScheduleType: dto.Prelim,
			}
			waveDB.Schedules = append(waveDB.Schedules, schedule)
		}

		if wave.DeltaTime != "" {
			schedule := Schedule{
				ScheduleTime: wave.DeltaTime,
				ScheduleType: dto.Delta,
			}
			waveDB.Schedules = append(waveDB.Schedules, schedule)
		}

		wpDB.Waves = append(wpDB.Waves, waveDB)
	}

	return wpDB
}

func TriggersFromSpannerToResponseView(spannerTriggers []TriggerItem) dto.GetTriggersResponse {
	triggers := make([]dto.TriggerResponse, 0)

	for _, st := range spannerTriggers {
		trigger := dto.TriggerResponse{
			RetailerID:     st.RetailerID,
			MfcID:          st.MfcID,
			Timezone:       st.Timezone,
			wpID:     st.wpID,
			WaveID:         st.WaveID,
			Cutoff:         st.Cutoff,
			ScheduleID:     st.ScheduleID,
			ScheduleTime:   st.ScheduleTime,
			TriggerAt:      st.TriggerAt,
			CreatedAt:      st.CreatedAt,
			CutoffDatetime: st.CutoffDateTime,
			FiredAt:        st.FiredAt,
		}

		triggers = append(triggers, trigger)
	}

	return dto.GetTriggersResponse{Triggers: triggers}
}

func FiredTriggersFromSpannerToResponseView(spannerTriggers []TriggerItem) dto.FireTriggersResponse {
	triggers := make([]dto.FireTriggerResponse, len(spannerTriggers))

	for idx, st := range spannerTriggers {
		trigger := dto.FireTriggerResponse{
			RetailerID:     st.RetailerID,
			MfcID:          st.MfcID,
			ScheduleID:     st.ScheduleID,
			TriggerAt:      st.TriggerAt,
			CreatedAt:      st.CreatedAt,
			FiredAt:        st.FiredAt,
			CutoffDateTime: st.CutoffDateTime,
		}

		triggers[idx] = trigger
	}

	return dto.FireTriggersResponse{Triggers: triggers}
}
