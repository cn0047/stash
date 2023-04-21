package repository

import (
	"strings"

	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/to-com/wp/config"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"time"
)

type Repository struct {
	logger   *zap.SugaredLogger
	dbClient *spanner.Client
}

func New(cfg *config.Config, logger *zap.SugaredLogger) (*Repository, error) {
	db := fmt.Sprintf(
		"projects/%s/instances/%s/databases/%s", cfg.ProjectID, cfg.SpannerInstance, cfg.SpannerDB,
	)
	dbClient, err := spanner.NewClient(context.Background(), db)
	if err != nil {
		return nil, err
	}
	logger.Debugf("initialized DB: %s", db)

	return &Repository{logger: logger, dbClient: dbClient}, nil
}

func (r *Repository) getwpID(ctx context.Context, txn *spanner.ReadWriteTransaction,
	retailerID, mfcID string) (string, error) {
	var wpID string

	stmt := spanner.NewStatement(`
						  SELECT wave_plan_id 
	                      FROM wave_plan
	                      WHERE mfc_id = @mfc_id AND retailer_id = @retailer_id`)

	stmt.Params["retailer_id"] = retailerID
	stmt.Params["mfc_id"] = mfcID

	iter := txn.Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return wpID, err
		}
		if err := row.Columns(&wpID); err != nil {
			return wpID, err
		}
	}
	return wpID, nil
}

func (r *Repository) Getwp(ctx context.Context, retailerID, mfcID string) ([]wpItem, error) {
	ctx, span := trace.StartSpan(ctx, "repository.get-wave-plan")
	defer span.End()

	stmt := spanner.NewStatement(
		`SELECT wp.wave_plan_id,
					 wp.retailer_id,
					 wp.mfc_id,
					 wp.timezone,
					 wp.created_at,
					 wp.created_by,
					 w.wave_id,
					 w.cutoff_time,
					 w.from_time,
					 w.to_time,
					 s.schedule_id,
					 s.schedule_type,
					 s.schedule_time
			  FROM wave_plan wp
			  JOIN wave w ON w.wave_plan_id = wp.wave_plan_id
			  LEFT JOIN schedule s ON s.wave_id = w.wave_id
			  WHERE wp.retailer_id = @retailer_id
			    AND wp.mfc_id = @mfc_id
			  ORDER BY w.cutoff_time`)

	stmt.Params["retailer_id"] = retailerID
	stmt.Params["mfc_id"] = mfcID

	r.logger.Infof("Retrieving wp for retailer %s, mfc: %s", retailerID, mfcID)

	iter := r.dbClient.Single().
		WithTimestampBound(spanner.MaxStaleness(15*time.Second)).
		Query(ctx, stmt)
	defer iter.Stop()

	var wp []wpItem

	for {
		row, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return wp, err
		}

		wpItem := wpItem{}
		if err := row.ToStruct(&wpItem); err != nil {
			return wp, err
		}

		wp = append(wp, wpItem)
	}

	r.logger.Infow("wp retrieved successfully",
		"retailer", retailerID, "mfc", mfcID, "wp", wp)

	return wp, nil
}

// Createwp removes, if exists, wave plan for combination of retailerID + mfcID,
// and creates a new one
func (r *Repository) Createwp(ctx context.Context, wp wp) (wp, error) {
	ctx, span := trace.StartSpan(ctx, "repository.create-wave-plan")
	defer span.End()

	retailerID := wp.RetailerID
	mfcID := wp.MfcID

	// since tables that contain wave_plan waves, schedules and triggers hierarchically interleave with each other,
	// we need to delete only one record in the "most parent" table, related records in child tables automatically deleted
	// "on cascade"; to use cascade delete we need to get the key ("wave_plan_id" field) of current wave plan

	_, err := r.dbClient.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		wpID, err := r.getwpID(ctx, txn, retailerID, mfcID)
		if err != nil {
			return err
		}

		r.logger.Infow("Creating new wave plan, active one will be deleted",
			"retailer", retailerID, "mfc", mfcID, "currentwpID", wpID, "wp", wp)
		// store all spanner mutations to be executed within a transaction
		var mutations []*spanner.Mutation
		var mutation *spanner.Mutation

		// single delete from parent table, this triggers cascade delete in child tables
		mutation = spanner.Delete("wave_plan", spanner.Key{wpID})
		mutations = append(mutations, mutation)
		wp.ID = uuid.NewString()
		wp.CreatedAt = time.Now().UTC()
		mutation, err = spanner.InsertStruct("wave_plan", wp)
		if err != nil {
			return err
		}
		mutations = append(mutations, mutation)

		for i := range wp.Waves {
			wp.Waves[i].ID = uuid.NewString()
			wp.Waves[i].wpID = wp.ID
			mutation, err = spanner.InsertStruct("wave", wp.Waves[i])
			if err != nil {
				return err
			}
			mutations = append(mutations, mutation)

			for j := range wp.Waves[i].Schedules {
				wp.Waves[i].Schedules[j].ID = uuid.NewString()
				wp.Waves[i].Schedules[j].WaveID = wp.Waves[i].ID
				wp.Waves[i].Schedules[j].wpID = wp.ID
				mutation, err = spanner.InsertStruct("schedule", wp.Waves[i].Schedules[j])
				if err != nil {
					return err
				}
				mutations = append(mutations, mutation)
			}
		}

		if err := txn.BufferWrite(mutations); err != nil {
			r.logger.Errorw("error occurred while creating wp",
				"retailer", wp.RetailerID, "mfc", wp.MfcID, zap.Error(err))

			return err
		}
		return nil
	})

	if err != nil {
		return wp, err
	}

	r.logger.Infow("wp successfully created", "retailer", wp.RetailerID, "mfc", wp.MfcID)

	return wp, nil
}

// CreateTriggers performs batch triggers insert in single transaction.
func (r *Repository) CreateTriggers(ctx context.Context, triggers []Trigger) (insertedCount int, err error) {
	ctx, span := trace.StartSpan(ctx, "repository.create-triggers")
	defer span.End()

	_, err = r.dbClient.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		for _, trigger := range triggers {
			trigger.CreatedAt = time.Now().UTC()
			mutation, err := spanner.InsertStruct("trigger", trigger)
			if err != nil {
				return fmt.Errorf("failed to create insert mutation, err: %w", err)
			}
			err = tx.BufferWrite([]*spanner.Mutation{mutation})
			if err == nil {
				insertedCount++
			} else {
				if strings.Contains(err.Error(), "AlreadyExists") {
					r.logger.Infof("trigger insertion skipped, trigger already exists: %+v", trigger)
				} else {
					return fmt.Errorf("failed to insert trigger, err: %w", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "AlreadyExists") {
			r.logger.Infof("already exists error appeared during transaction execution: %+v", err)
		} else {
			return insertedCount, fmt.Errorf("failed to execute transaction, err: %w", err)
		}
	}

	return insertedCount, nil
}

func (r *Repository) GetScheduleItems(ctx context.Context) ([]ScheduleItem, error) {
	ctx, span := trace.StartSpan(ctx, "repository.get-schedule-items")
	defer span.End()

	stmt := spanner.NewStatement(
		`SELECT wp.retailer_id,
					wp.mfc_id,
					wp.timezone,
					s.wave_plan_id,
					s.wave_id,
					s.schedule_id,
					s.schedule_time,
					w.cutoff_time,
					t.trigger_at,
					t.fired_at
			FROM schedule s
			JOIN wave_plan wp ON s.wave_plan_id = wp.wave_plan_id
			JOIN wave w ON s.wave_id = w.wave_id
			LEFT JOIN trigger t ON t.schedule_id = s.schedule_id AND t.fired_at IS NULL
			WHERE t.trigger_at IS NULL`)

	r.logger.Infof("Retrieving all schedules")

	var schedules []ScheduleItem

	iter := r.dbClient.Single().Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return schedules, err
		}

		scheduleItem := ScheduleItem{}
		if err := row.ToStruct(&scheduleItem); err != nil {
			return schedules, err
		}

		schedules = append(schedules, scheduleItem)
	}

	return schedules, nil
}

func (r *Repository) GetTriggers(ctx context.Context, retailerID, mfcID string) ([]TriggerItem, error) {
	ctx, span := trace.StartSpan(ctx, "repository.get-triggers")
	defer span.End()

	sql := `SELECT wp.retailer_id,
				   wp.mfc_id,
                   wp.timezone,
				   wp.wave_plan_id,
                   w.wave_id,
                   w.cutoff_time,
                   s.schedule_id,
                   s.schedule_time,
                   t.trigger_at,
                   t.created_at,
                   t.cutoff_datetime,
                   t.fired_at
            FROM wave_plan wp
            JOIN wave w ON w.wave_plan_id = wp.wave_plan_id
            JOIN schedule s ON s.wave_plan_id = w.wave_plan_id
              AND s.wave_id = w.wave_id
            JOIN trigger t ON t.wave_plan_id = s.wave_plan_id
              AND t.wave_id = s.wave_id
              AND t.schedule_id = s.schedule_id
            WHERE wp.retailer_id = @retailer_id
             AND wp.mfc_id = @mfc_id
             AND t.trigger_at  >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 DAY)
             ORDER BY t.trigger_at ASC LIMIT 100`

	stmt := spanner.NewStatement(sql)
	stmt.Params["retailer_id"] = retailerID
	stmt.Params["mfc_id"] = mfcID

	r.logger.Infof("Retrieving triggers for retailer %s, mfc: %s", retailerID, mfcID)

	iter := r.dbClient.Single().Query(ctx, stmt)
	defer iter.Stop()

	var triggers []TriggerItem

	for {
		row, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return triggers, err
		}

		triggerItem := TriggerItem{}
		if err := row.ToStruct(&triggerItem); err != nil {
			return triggers, err
		}

		triggers = append(triggers, triggerItem)
	}

	r.logger.Infow("triggers retrieved successfully", "retailer", retailerID, "mfc", mfcID)

	return triggers, nil
}

func (r *Repository) SelectTriggersToFire(ctx context.Context) ([]TriggerItem, error) {
	stmt := spanner.NewStatement(`SELECT wp.retailer_id,
	                                    wp.wave_plan_id,
	                                    wp.mfc_id,
	                                    s.schedule_id,
	                                    s.schedule_type,
	                                    s.wave_id,
	                                    t.cutoff_datetime,
	                                    t.trigger_at,
	                                    t.created_at,
	                                    t.fired_at
                                 FROM trigger t
                                 JOIN wave_plan wp on wp.wave_plan_id = t.wave_plan_id
                                 JOIN schedule s ON s.schedule_id = t.schedule_id
                                 WHERE t.fired_at is NULL 
                                 AND t.trigger_at < CURRENT_TIMESTAMP`)

	iter := r.dbClient.Single().Query(ctx, stmt)
	defer iter.Stop()
	var triggers []TriggerItem

	for {
		row, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return triggers, err
		}

		triggerItem := TriggerItem{}
		if err := row.ToStruct(&triggerItem); err != nil {
			return triggers, err
		}

		triggers = append(triggers, triggerItem)
	}
	if len(triggers) > 0 {
		r.logger.Infow("selected triggers to be fired", "triggers", triggers)
	} else {
		r.logger.Info("No triggers to be fired")
	}

	return triggers, nil
}

func (r *Repository) MarkTriggersAsFired(ctx context.Context, triggers []TriggerItem, firedAt time.Time) error {
	_, errTxn := r.dbClient.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		for _, trg := range triggers {
			stmt := spanner.NewStatement(`UPDATE trigger SET fired_at = @fired_at
               									WHERE wave_plan_id = @wave_plan_id
               									AND wave_id        = @wave_id
    											AND schedule_id    = @schedule_id
               									AND trigger_at     = @trigger_at`)

			stmt.Params["fired_at"] = firedAt
			stmt.Params["wave_plan_id"] = trg.wpID
			stmt.Params["wave_id"] = trg.WaveID
			stmt.Params["schedule_id"] = trg.ScheduleID
			stmt.Params["trigger_at"] = trg.TriggerAt

			if _, errUpdate := txn.Update(ctx, stmt); errUpdate != nil {
				return errUpdate
			}
			r.logger.Infow("Trigger marked as fired", "wpID", trg.wpID,
				"waveID", trg.WaveID, "scheduleID", trg.ScheduleID, "triggerAt", trg.TriggerAt)
		}
		return nil
	})
	return errTxn
}
