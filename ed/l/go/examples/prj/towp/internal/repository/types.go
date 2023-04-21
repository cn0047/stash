package repository

import (
	"time"

	"cloud.google.com/go/spanner"
)

/*
	This file contains DB related data models
*/

type wp struct {
	ID         string    `spanner:"wave_plan_id"`
	RetailerID string    `spanner:"retailer_id"`
	MfcID      string    `spanner:"mfc_id"`
	Timezone   string    `spanner:"timezone"`
	CreatedAt  time.Time `spanner:"created_at"`
	CreatedBy  string    `spanner:"created_by"`
	Waves      []Wave    `spanner:"-"`
}

type Wave struct {
	wpID string     `spanner:"wave_plan_id"`
	ID         string     `spanner:"wave_id"`
	Cutoff     string     `spanner:"cutoff_time"`
	FromTime   string     `spanner:"from_time"`
	ToTime     string     `spanner:"to_time"`
	Schedules  []Schedule `spanner:"-"`
}

type Schedule struct {
	wpID   string `spanner:"wave_plan_id"`
	WaveID       string `spanner:"wave_id"`
	ID           string `spanner:"schedule_id"`
	ScheduleType string `spanner:"schedule_type"`
	ScheduleTime string `spanner:"schedule_time"`
}

type wpItem struct {
	wpID   string             `spanner:"wave_plan_id"`
	RetailerID   string             `spanner:"retailer_id"`
	MfcID        string             `spanner:"mfc_id"`
	Timezone     string             `spanner:"timezone"`
	CreatedAt    time.Time          `spanner:"created_at"`
	CreatedBy    string             `spanner:"created_by"`
	WaveID       string             `spanner:"wave_id"`
	Cutoff       string             `spanner:"cutoff_time"`
	FromTime     string             `spanner:"from_time"`
	ToTime       string             `spanner:"to_time"`
	ScheduleID   spanner.NullString `spanner:"schedule_id"`
	ScheduleType spanner.NullString `spanner:"schedule_type"`
	ScheduleTime spanner.NullString `spanner:"schedule_time"`
}

type ScheduleItem struct {
	RetailerID   string     `spanner:"retailer_id"`
	MfcID        string     `spanner:"mfc_id"`
	Timezone     string     `spanner:"timezone"`
	wpID   string     `spanner:"wave_plan_id"`
	WaveID       string     `spanner:"wave_id"`
	ScheduleID   string     `spanner:"schedule_id"`
	ScheduleTime string     `spanner:"schedule_time"`
	CutoffTime   string     `spanner:"cutoff_time"`
	TriggerAt    *time.Time `spanner:"trigger_at"`
	FiredAt      *time.Time `spanner:"fired_at"`
}

type Trigger struct {
	wpID     string     `spanner:"wave_plan_id"`
	WaveID         string     `spanner:"wave_id"`
	ScheduleID     string     `spanner:"schedule_id"`
	TriggerAt      time.Time  `spanner:"trigger_at"`
	CutoffDateTime time.Time  `spanner:"cutoff_datetime"`
	CreatedAt      time.Time  `spanner:"created_at"`
	FiredAt        *time.Time `spanner:"fired_at"`
}

type TriggerItem struct {
	RetailerID     string     `spanner:"retailer_id"`
	MfcID          string     `spanner:"mfc_id"`
	Timezone       string     `spanner:"timezone"`
	wpID     string     `spanner:"wave_plan_id"`
	WaveID         string     `spanner:"wave_id"`
	Cutoff         string     `spanner:"cutoff_time"`
	ScheduleID     string     `spanner:"schedule_id"`
	ScheduleTime   string     `spanner:"schedule_time"`
	ScheduleType   string     `spanner:"schedule_type"`
	TriggerAt      time.Time  `spanner:"trigger_at"`
	CreatedAt      time.Time  `spanner:"created_at"`
	CutoffDateTime time.Time  `spanner:"cutoff_datetime"`
	FiredAt        *time.Time `spanner:"fired_at"`
}
