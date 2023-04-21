package dto

import (
	"time"
)

/*
	This file contains data models that reflect the service endpoints requests/responses
*/

type wpRequest struct {
	RetailerID string        `json:"-"`
	MfcID      string        `json:"-"`
	UserID     string        `json:"-"`
	Waves      []WaveRequest `json:"waves"`
}

type WaveRequest struct {
	Cutoff     string `json:"cutoff_time"`
	FromTime   string `json:"from_time"`
	ToTime     string `json:"to_time"`
	PrelimTime string `json:"prelim_picklist_schedule_time"`
	DeltaTime  string `json:"delta_picklist_schedule_time"`
}

type wpResponse struct {
	ID         string         `json:"id"`
	RetailerID string         `json:"retailer_id"`
	MfcID      string         `json:"mfc_id"`
	Timezone   string         `json:"timezone"`
	CreatedAt  time.Time      `json:"created_time"`
	CreatedBy  string         `json:"created_by"`
	Waves      []WaveResponse `json:"waves"`
}

type WaveResponse struct {
	ID        string             `json:"id"`
	Cutoff    string             `json:"cutoff_time"`
	FromTime  string             `json:"from_time"`
	ToTime    string             `json:"to_time"`
	Schedules []ScheduleResponse `spanner:"-" json:"schedules"`
}

const (
	Prelim string = "prelim_picklist"
	Delta  string = "delta_picklist"
)

type ScheduleResponse struct {
	ID           string `json:"id"`
	ScheduleType string `json:"schedule_type"`
	ScheduleTime string `json:"schedule_time"`
}

type wpCreatedEvent struct {
	CreatedAt time.Time        `json:"created_time"`
	wp  wpResponse `json:"wave_plan"`
}

type GenerateTriggersResponse struct {
	GeneratedTriggers int `json:"generated_triggers"`
}

type TriggerResponse struct {
	RetailerID     string     `json:"retailer_id"`
	MfcID          string     `json:"mfc_id"`
	Timezone       string     `json:"timezone"`
	wpID     string     `json:"wave_plan_id"`
	WaveID         string     `json:"wave_id"`
	Cutoff         string     `json:"cutoff_time"`
	ScheduleID     string     `json:"schedule_id"`
	ScheduleTime   string     `json:"schedule_time"`
	TriggerAt      time.Time  `json:"trigger_at"`
	CreatedAt      time.Time  `json:"created_at"`
	CutoffDatetime time.Time  `json:"cutoff_datetime"`
	FiredAt        *time.Time `json:"fired_at"`
}

type GetTriggersResponse struct {
	Triggers []TriggerResponse `json:"triggers"`
}

type FireTriggerResponse struct {
	RetailerID     string     `json:"retailer_id"`
	MfcID          string     `json:"mfc_id"`
	ScheduleID     string     `json:"schedule_id"`
	ScheduleType   string     `json:"schedule_type"`
	CutoffDateTime time.Time  `json:"cutoff_datetime"`
	TriggerAt      time.Time  `json:"trigger_at"`
	CreatedAt      time.Time  `json:"created_at"`
	FiredAt        *time.Time `json:"fired_at"`
}

type FireTriggersResponse struct {
	Triggers []FireTriggerResponse `json:"triggers"`
}

type TriggersFiredEvent struct {
	ScheduleType string      `json:"schedule_type"`
	Cutoffs      []time.Time `json:"cutoffs"`
}
