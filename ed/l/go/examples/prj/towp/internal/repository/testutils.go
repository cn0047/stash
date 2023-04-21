package repository

import (
	"cloud.google.com/go/spanner"
	"github.com/to-com/wp/internal/dto"
	"time"
)

const AsiaDubai = "Asia/Dubai"

var nowUTC = func() time.Time {
	return time.Now().UTC()
}

var (
	maf            = "maf"
	D02            = "D02"
	D03            = "D03"
	wpNoIsps = wp{
		RetailerID: maf,
		MfcID:      D02,
		CreatedAt:  nowUTC(),
		CreatedBy:  "user-1",
		Timezone:   AsiaDubai,
		Waves: []Wave{
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
	wpWithIsps = wp{
		RetailerID: maf,
		MfcID:      D03,
		CreatedAt:  nowUTC(),
		CreatedBy:  "user-2",
		Timezone:   AsiaDubai,
		Waves: []Wave{
			{
				Cutoff:   "08:00",
				FromTime: "09:00",
				ToTime:   "11:59",
				Schedules: []Schedule{
					{
						ScheduleType: dto.Prelim,
						ScheduleTime: "04:00",
					},
					{
						ScheduleType: dto.Delta,
						ScheduleTime: "07:00",
					},
				},
			},
			{
				Cutoff:   "10:00",
				FromTime: "12:00",
				ToTime:   "16:59",
				Schedules: []Schedule{
					{
						ScheduleType: dto.Prelim,
						ScheduleTime: "04:00",
					},
				},
			},
			{
				Cutoff:   "15:00",
				FromTime: "17:00",
				ToTime:   "08:59",
				Schedules: []Schedule{
					{
						ScheduleType: dto.Delta,
						ScheduleTime: "14:00",
					},
				},
			},
		},
	}

	wpWithOneSchedule = wp{
		RetailerID: maf,
		MfcID:      D03,
		CreatedAt:  nowUTC(),
		CreatedBy:  "user-2",
		Timezone:   AsiaDubai,
		Waves: []Wave{
			{
				Cutoff:   "10:00",
				FromTime: "12:00",
				ToTime:   "14:59",
				Schedules: []Schedule{
					{
						ScheduleType: dto.Prelim,
						ScheduleTime: "09:00",
					},
				},
			},
			{
				Cutoff:   "13:00",
				FromTime: "15:00",
				ToTime:   "11:59",
			},
		},
	}
)

func makeTriggers(plan wp, triggerAt time.Time) []Trigger {
	var triggers []Trigger
	var wpID = plan.ID

	for _, wave := range plan.Waves {
		for _, schedule := range wave.Schedules {
			trg := Trigger{
				wpID:     wpID,
				WaveID:         wave.ID,
				ScheduleID:     schedule.ID,
				TriggerAt:      triggerAt,
				CutoffDateTime: nowUTC().AddDate(0, 0, -1),
			}
			triggers = append(triggers, trg)
		}
	}
	return triggers
}

// triggersToTriggerItems transforms spanner Trigger structs
// into business-level structs TriggerItem with only fields
// containing primary key to lookup triggers in database
func triggersToTriggerItems(triggers []Trigger) []TriggerItem {
	var triggerItems []TriggerItem
	for _, trg := range triggers {
		trgItem := TriggerItem{
			RetailerID: maf,
			MfcID:      D02,
			wpID: trg.wpID,
			WaveID:     trg.WaveID,
			ScheduleID: trg.ScheduleID,
			TriggerAt:  trg.TriggerAt,
		}
		triggerItems = append(triggerItems, trgItem)
	}
	return triggerItems
}

// triggerItemsToTriggers transforms business-level TriggerItem structs
// into API response structs TriggerItem
func triggerItemsToTriggers(triggerItems []TriggerItem) []Trigger {
	var triggers []Trigger
	for _, trgItem := range triggerItems {
		trg := Trigger{
			wpID:     trgItem.wpID,
			WaveID:         trgItem.WaveID,
			ScheduleID:     trgItem.ScheduleID,
			TriggerAt:      trgItem.TriggerAt,
			CutoffDateTime: trgItem.CutoffDateTime,
			CreatedAt:      time.Time{},
			FiredAt:        trgItem.FiredAt,
		}
		triggers = append(triggers, trg)
	}
	return triggers
}

func getIDs(wp wp) (string, []string, []string) {
	wavesIDs := make([]string, 0, len(wp.Waves))
	schedulesIDs := make([]string, 0)
	for i := range wp.Waves {
		wavesIDs = append(wavesIDs, wp.Waves[i].ID)
		for j := range wp.Waves[i].Schedules {
			schedulesIDs = append(schedulesIDs, wp.Waves[i].Schedules[j].ID)
		}
	}

	return wp.ID, wavesIDs, schedulesIDs
}

func towpItems(wp wp) []wpItem {
	var wpItems []wpItem
	for _, wave := range wp.Waves {
		if wave.Schedules != nil {
			for _, schedule := range wave.Schedules {
				wpItems = append(wpItems, wpItem{
					wpID: wp.ID,
					RetailerID: wp.RetailerID,
					MfcID:      wp.MfcID,
					CreatedAt:  wp.CreatedAt,
					CreatedBy:  wp.CreatedBy,
					WaveID:     wave.ID,
					Cutoff:     wave.Cutoff,
					FromTime:   wave.FromTime,
					ToTime:     wave.ToTime,
					Timezone:   wp.Timezone,
					ScheduleID: spanner.NullString{
						StringVal: schedule.ID,
						Valid:     true,
					},
					ScheduleType: spanner.NullString{
						StringVal: schedule.ScheduleType,
						Valid:     true,
					},
					ScheduleTime: spanner.NullString{
						StringVal: schedule.ScheduleTime,
						Valid:     true,
					},
				})
			}
		} else {
			wpItems = append(wpItems, wpItem{
				wpID: wp.ID,
				RetailerID: wp.RetailerID,
				MfcID:      wp.MfcID,
				CreatedAt:  wp.CreatedAt,
				CreatedBy:  wp.CreatedBy,
				WaveID:     wave.ID,
				Cutoff:     wave.Cutoff,
				FromTime:   wave.FromTime,
				ToTime:     wave.ToTime,
				Timezone:   wp.Timezone,
			})
		}
	}

	return wpItems
}
