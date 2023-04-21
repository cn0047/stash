package repository

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"google.golang.org/api/iterator"
	"testing"
	"time"
)

type RepositorySuite struct {
	suite.Suite
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

var (
	rep  *Repository
	logs *observer.ObservedLogs
)

func (s *RepositorySuite) newRep() (*Repository, *observer.ObservedLogs, error) {
	cfg, err := config.Load()
	if err != nil {
		s.T().Errorf("unable to load config for tests, error: %v", err)
		s.T().FailNow()
	}
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()

	rep, err := New(cfg, observedLogger)
	if err != nil {
		s.T().Errorf("unable to initialise repository for tests, error: %v", err)
		s.T().FailNow()
	}
	return rep, observedLogs, nil
}

func getTestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")
	return ctx
}

func (s *RepositorySuite) SetupSuite() {
	s.T().Setenv("SPANNER_EMULATOR_HOST", "localhost:9010")
}

func (s *RepositorySuite) SetupTest() {
	rep, logs, _ = s.newRep()
	s.cleanUpDB(getTestContext())
}

func (s *RepositorySuite) TearDownSuite() {
	s.cleanUpDB(getTestContext())
}

func (s *RepositorySuite) getIDListFromTable(ctx context.Context, tableName string, columnName string) ([]string, error) {
	stmt := spanner.NewStatement(fmt.Sprintf("select %s from %s", columnName, tableName))
	iter := rep.dbClient.Single().Query(ctx, stmt)
	defer iter.Stop()

	var ids []string
	var id string

	for {
		row, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		if err = row.Column(0, &id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *RepositorySuite) readTriggers(ctx context.Context) ([]Trigger, error) {
	stmt := spanner.NewStatement("select * from trigger")
	iter := rep.dbClient.Single().Query(ctx, stmt)
	defer iter.Stop()

	var triggers []Trigger
	var trg Trigger

	for {
		row, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}

		if err = row.ToStruct(&trg); err != nil {
			return nil, err
		}
		triggers = append(triggers, trg)
	}
	return triggers, nil
}

func (s *RepositorySuite) cleanUpDB(ctx context.Context) {
	var mutations []*spanner.Mutation
	var mutation *spanner.Mutation
	ids, err := s.getIDListFromTable(ctx, "wave_plan", "wave_plan_id")
	if err != nil {
		s.T().Errorf("fail to get all wave plans ids, %v", err)
		s.T().FailNow()
	}
	for _, id := range ids {
		mutation = spanner.Delete("wave_plan", spanner.Key{id})
		mutations = append(mutations, mutation)
	}

	_, err = rep.dbClient.Apply(ctx, mutations)
	if err != nil {
		s.T().Errorf("fail to remove data from wave planner DB, %v", err)
		s.T().FailNow()
	}

	s.T().Log("db was cleaned up")
}

func (s *RepositorySuite) validateDBState(ctx context.Context, wpIDs []string, wavesIDs []string, schedulesIDs []string) {
	ids, err := s.getIDListFromTable(ctx, "wave_plan", "wave_plan_id")
	s.Nil(err)
	s.ElementsMatch(ids, wpIDs)

	ids, err = s.getIDListFromTable(ctx, "wave", "wave_id")
	s.Nil(err)
	s.ElementsMatch(ids, wavesIDs)

	ids, err = s.getIDListFromTable(ctx, "schedule", "schedule_id")
	s.Nil(err)
	s.ElementsMatch(ids, schedulesIDs)
}

func (s *RepositorySuite) verifyNotEmptyData(wp wp) {
	s.NotEmpty(wp.ID)
	s.NotEmpty(wp.CreatedAt)
	for _, wave := range wp.Waves {
		s.NotEmpty(wave.ID)
		s.NotEmpty(wave.wpID)
		s.Equal(wp.ID, wave.wpID)
		for _, schedule := range wave.Schedules {
			s.NotEmpty(schedule.ID)
			s.NotEmpty(schedule.WaveID)
			s.Equal(wave.ID, schedule.WaveID)
			s.NotEmpty(schedule.wpID)
			s.Equal(wp.ID, schedule.wpID)
		}
	}
}

func (s *RepositorySuite) TestCreatewpWithNoSchedules() {
	ctx := getTestContext()

	wp1 := wpNoIsps

	s.T().Log("create wave plan with no schedulers")
	wp1, err := rep.Createwp(ctx, wp1)
	s.Nil(err)
	s.verifyNotEmptyData(wp1)

	wpID, waves, schedules := getIDs(wp1)
	s.T().Log("validate wave plan in DB")
	s.validateDBState(ctx, []string{wpID}, waves, schedules)
}

func (s *RepositorySuite) TestCreatewpWithSchedules() {
	ctx := getTestContext()

	wp := wpWithIsps
	s.T().Log("create wave plan with schedulers")
	wp, err := rep.Createwp(ctx, wp)
	s.Nil(err)
	s.verifyNotEmptyData(wp)

	wpID, waves, schedules := getIDs(wp)
	s.T().Log("validate wave plan in DB")
	s.validateDBState(ctx, []string{wpID}, waves, schedules)
}

func (s *RepositorySuite) TestCreateTwowpsForTheSameRetailerAndMfc() {
	ctx := getTestContext()

	wp1 := wpNoIsps
	wp2 := wpNoIsps
	s.T().Log("create one wave plan with no schedulers")
	wp1, err := rep.Createwp(ctx, wp1)
	s.Nil(err)

	wpID, waves, schedules := getIDs(wp1)
	s.T().Log("validate wave plan in DB")
	s.validateDBState(ctx, []string{wpID}, waves, schedules)

	s.T().Log("create the other wave plan for the same retailer and mfc")
	wp2, err = rep.Createwp(ctx, wp2)
	s.Nil(err)

	wpID, waves, schedules = getIDs(wp2)
	s.T().Log("validate wave plan in DB")
	s.validateDBState(ctx, []string{wpID}, waves, schedules)
}

func (s *RepositorySuite) TestCreateTwowpsForTheSameRetailerDifferentMfcs() {
	ctx := getTestContext()

	wp1 := wpNoIsps
	wp2 := wpWithIsps
	wp1.MfcID = D02
	wp2.MfcID = D03

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp1.RetailerID, wp1.MfcID)
	wp1, err := rep.Createwp(ctx, wp1)
	s.Nil(err)

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp2.RetailerID, wp2.MfcID)
	wp2, err = rep.Createwp(ctx, wp2)
	s.Nil(err)

	wp1ID, waves1, schedules1 := getIDs(wp1)
	wp2ID, waves2, schedules2 := getIDs(wp2)
	s.T().Log("validate wave plans in DB")
	s.validateDBState(ctx, []string{wp1ID, wp2ID}, append(waves1, waves2...), append(schedules1, schedules2...))
}

func (s *RepositorySuite) TestCreateTwowpsForDifferentRetailersDifferentMfcs() {
	ctx := getTestContext()

	wp1 := wpNoIsps
	wp2 := wpWithIsps
	wp1.MfcID = D02
	wp1.RetailerID = maf
	wp2.MfcID = "3116"
	wp2.RetailerID = "abs"

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp1.RetailerID, wp1.MfcID)
	wp1, err := rep.Createwp(ctx, wp1)
	s.Nil(err)

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp2.RetailerID, wp2.MfcID)
	wp2, err = rep.Createwp(ctx, wp2)
	s.Nil(err)

	wp1ID, waves1, schedules1 := getIDs(wp1)
	wp2ID, waves2, schedules2 := getIDs(wp2)
	s.T().Log("validate wave plans in DB")
	s.validateDBState(ctx, []string{wp1ID, wp2ID}, append(waves1, waves2...), append(schedules1, schedules2...))
}

func (s *RepositorySuite) TestGetwp() {
	ctx := getTestContext()

	wp1 := wpNoIsps
	wp2 := wpWithIsps
	wp3 := wpNoIsps
	wp1.MfcID = D02
	wp1.RetailerID = maf
	wp2.MfcID = D03
	wp2.RetailerID = maf
	wp3.MfcID = "3116"
	wp3.RetailerID = "abs"

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp1.RetailerID, wp1.MfcID)
	wp1, err := rep.Createwp(ctx, wp1)
	s.Nil(err)

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp2.RetailerID, wp2.MfcID)
	wp2, err = rep.Createwp(ctx, wp2)
	s.Nil(err)

	s.T().Logf("create wave plan for retailer %v and mfc %v", wp3.RetailerID, wp3.MfcID)
	wp3, err = rep.Createwp(ctx, wp3)
	s.Nil(err)

	s.T().Log("get nonexistent wave plan")
	wpItems, err := rep.Getwp(ctx, wp2.RetailerID, wp3.MfcID)
	s.Nil(err)
	s.Empty(wpItems)

	s.T().Log("get existing wave plan")
	wpItems, err = rep.Getwp(ctx, wp2.RetailerID, wp2.MfcID)
	s.Nil(err)

	s.T().Log("verify data stored to DB")
	s.verifyNotEmptyData(wp2)
	expectedwpItems := towpItems(wp2)
	s.ElementsMatch(wpItems, expectedwpItems)

	allLogs := logs.All()
	s.Equal(10, len(allLogs))
	s.Equal(allLogs[0].Message, "Creating new wave plan, active one will be deleted")
	s.Equal(allLogs[1].Message, "wp successfully created")
	s.Equal(allLogs[2].Message, "Creating new wave plan, active one will be deleted")
	s.Equal(allLogs[3].Message, "wp successfully created")
	s.Equal(allLogs[4].Message, "Creating new wave plan, active one will be deleted")
	s.Equal(allLogs[5].Message, "wp successfully created")
	s.Equal(allLogs[6].Message, fmt.Sprintf("Retrieving wp for retailer %s, mfc: %s", wp2.RetailerID, wp3.MfcID))
	s.Equal(allLogs[7].Message, "wp retrieved successfully")
	s.Equal(allLogs[8].Message, fmt.Sprintf("Retrieving wp for retailer %s, mfc: %s", wp2.RetailerID, wp2.MfcID))
	s.Equal(allLogs[9].Message, "wp retrieved successfully")
}

func (s *RepositorySuite) TestSelectTriggersToFire() {
	ctx := getTestContext()
	wp1 := wpWithIsps

	s.T().Log("create wave plan")
	wp1, _ = rep.Createwp(ctx, wp1)

	pastTriggerAt := nowUTC().AddDate(0, 0, -1)
	futureTriggerAt := nowUTC().AddDate(0, 0, 1)
	pastTriggers := makeTriggers(wp1, pastTriggerAt)
	// will create future triggers that should not be fired
	futureTriggers := makeTriggers(wp1, futureTriggerAt)

	triggers := append(pastTriggers, futureTriggers...)

	_, err := rep.CreateTriggers(ctx, triggers)
	s.Nil(err)

	triggerItemsToFire, err := rep.SelectTriggersToFire(ctx)
	triggersToFire := triggerItemsToTriggers(triggerItemsToFire)
	s.Nil(err)

	s.ElementsMatch(triggersToFire, pastTriggers)
}

func (s *RepositorySuite) TestFireTriggers() {
	ctx := getTestContext()
	wp1 := wpWithIsps

	s.T().Log("create wave plan")
	wp1, _ = rep.Createwp(ctx, wp1)

	triggerAt := nowUTC().AddDate(0, 0, -1)
	trgs := makeTriggers(wp1, triggerAt)

	triggerItems := triggersToTriggerItems(trgs)

	_, err := rep.CreateTriggers(ctx, trgs)
	s.Nil(err)

	firedAt := nowUTC()
	fireErr := rep.MarkTriggersAsFired(ctx, triggerItems, firedAt)
	s.Nil(fireErr)

	res, err := s.readTriggers(ctx)
	for _, tr := range res {
		s.Equal(&firedAt, tr.FiredAt)
	}
	s.Nil(err)
}

func (s *RepositorySuite) TestFindScheduleItem() {
	ctx := getTestContext()

	wp := wpWithOneSchedule
	wp, err := rep.Createwp(ctx, wp)
	s.Nil(err)

	scheduleItems, err := rep.GetScheduleItems(ctx)

	s.Nil(err)
	s.Equal(1, len(scheduleItems))
	s.Equal(ScheduleItem{
		RetailerID:   wp.RetailerID,
		MfcID:        wp.MfcID,
		Timezone:     wp.Timezone,
		wpID:   wp.ID,
		WaveID:       wp.Waves[0].ID,
		ScheduleID:   wp.Waves[0].Schedules[0].ID,
		ScheduleTime: wp.Waves[0].Schedules[0].ScheduleTime,
		CutoffTime:   wp.Waves[0].Cutoff,
		TriggerAt:    nil,
		FiredAt:      nil,
	}, scheduleItems[0])
	allLogs := logs.All()
	s.Equal(3, len(allLogs))
	s.Equal("Creating new wave plan, active one will be deleted", allLogs[0].Message)
	s.Equal("wp successfully created", allLogs[1].Message)
	s.Equal("Retrieving all schedules", allLogs[2].Message)
}

func (s *RepositorySuite) TestNoSchedules() {
	ctx := getTestContext()

	wp := wpNoIsps
	_, err := rep.Createwp(ctx, wp)
	s.Nil(err)

	scheduleItems, err := rep.GetScheduleItems(ctx)

	s.Nil(err)
	s.Equal(0, len(scheduleItems))
	allLogs := logs.All()
	s.Equal(3, len(allLogs))
	s.Equal("Creating new wave plan, active one will be deleted", allLogs[0].Message)
	s.Equal("wp successfully created", allLogs[1].Message)
}

func (s *RepositorySuite) TestMoreThanOneSchedule() {
	ctx := getTestContext()

	testwp := wpWithIsps
	_, err := rep.Createwp(ctx, testwp)
	s.Nil(err)

	scheduleItems, err := rep.GetScheduleItems(ctx)

	s.Nil(err)
	s.Equal(4, len(scheduleItems))
	allLogs := logs.All()
	s.Equal(3, len(allLogs))
	s.Equal("Creating new wave plan, active one will be deleted", allLogs[0].Message)
	s.Equal("wp successfully created", allLogs[1].Message)
}

func (s *RepositorySuite) TestCreateTriggers() {
	ctx := getTestContext()
	wp := wpWithIsps
	wp, err := rep.Createwp(ctx, wp)
	s.Nil(err)

	triggers := makeTriggers(wp, time.Now().Add(2*time.Hour).UTC())
	_, err = rep.CreateTriggers(ctx, triggers)
	s.Nil(err)

	triggerItems, err := rep.GetTriggers(ctx, wp.RetailerID, wp.MfcID)
	actualTriggers := triggerItemsToTriggers(triggerItems)
	s.Nil(err)
	s.Equal(4, len(actualTriggers))
	s.NotEmpty(triggerItems[0].CreatedAt)
	s.ElementsMatch(triggers, actualTriggers)
	allLogs := logs.All()
	s.Equal(4, len(allLogs))
	s.Equal("Creating new wave plan, active one will be deleted", allLogs[0].Message)
	s.Equal("wp successfully created", allLogs[1].Message)
	s.Equal(fmt.Sprintf("Retrieving triggers for retailer %s, mfc: %s", wp.RetailerID, wp.MfcID),
		allLogs[2].Message)
	s.Equal("triggers retrieved successfully", allLogs[3].Message)
}

func (s *RepositorySuite) TestCreateTriggersWithError() {
	ctx := getTestContext()
	wp, waveIDs, scheduleIDs := getIDs(wpWithOneSchedule)
	trigger1 := &Trigger{
		wpID:     wp,
		WaveID:         waveIDs[0],
		ScheduleID:     scheduleIDs[0],
		TriggerAt:      time.Now().Add(-1 * time.Hour),
		CutoffDateTime: time.Now().Add(-2 * time.Hour),
		FiredAt:        nil,
	}
	triggers := append(make([]Trigger, 0), *trigger1)
	_, err := rep.CreateTriggers(ctx, triggers)
	s.Contains(err.Error(), "Insert failed because key was not found in parent table")
}

func (s *RepositorySuite) TestGetTriggers() {
	ctx := getTestContext()
	wp := wpWithOneSchedule
	wp, err := rep.Createwp(ctx, wp)
	triggers := make([]Trigger, 0)
	s.Nil(err)

	triggers1 := makeTriggers(wp, time.Now().Add(-25*time.Hour).UTC())
	trigger2 := makeTriggers(wp, time.Now().Add(-2*time.Hour).UTC())[0]
	trigger3 := makeTriggers(wp, time.Now().UTC())[0]
	triggers = append(triggers, trigger2, trigger3)
	_, err = rep.CreateTriggers(ctx, triggers)
	s.Nil(err)
	_, err = rep.CreateTriggers(ctx, triggers1)
	s.Nil(err)

	triggerItems, err := rep.GetTriggers(ctx, wp.RetailerID, wp.MfcID)
	actualTriggers := triggerItemsToTriggers(triggerItems)
	s.Nil(err)
	s.Equal(2, len(actualTriggers))
	s.Equal(triggers, actualTriggers)
}

func (s *RepositorySuite) TestGetTriggersEmpty() {
	ctx := getTestContext()
	wp := wpWithIsps
	wp, err := rep.Createwp(ctx, wp)
	s.Nil(err)

	triggers := makeTriggers(wp, time.Now().Add(-25*time.Hour).UTC())
	_, err = rep.CreateTriggers(ctx, triggers)
	s.Nil(err)

	s.T().Run("no triggers within 24 hours period", func(t *testing.T) {
		triggerItems, err := rep.GetTriggers(ctx, wp.RetailerID, wp.MfcID)
		actualTriggers := triggerItemsToTriggers(triggerItems)
		s.Nil(err)
		s.Equal(0, len(actualTriggers))
		s.ElementsMatch(Trigger{}, actualTriggers)
	})

	triggers = makeTriggers(wp, time.Now().Add(-22*time.Hour).UTC())
	rep.CreateTriggers(ctx, triggers)
	s.T().Run("no triggers for fake retailer", func(t *testing.T) {
		triggerItems, err := rep.GetTriggers(ctx, "fake-retailer", wp.MfcID)
		actualTriggers := triggerItemsToTriggers(triggerItems)
		s.Nil(err)
		s.Equal(0, len(actualTriggers))
		s.ElementsMatch(Trigger{}, actualTriggers)
	})
	s.T().Run("no triggers for fake mfc", func(t *testing.T) {
		triggerItems, err := rep.GetTriggers(ctx, wp.RetailerID, "fake-mfc")
		actualTriggers := triggerItemsToTriggers(triggerItems)
		s.Nil(err)
		s.Equal(0, len(actualTriggers))
		s.ElementsMatch(Trigger{}, actualTriggers)
	})
}
