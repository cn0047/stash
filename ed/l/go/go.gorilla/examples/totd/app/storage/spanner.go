// Package storage represents package to work with GCP Spanner.
package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"

	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/app/util"
)

const (
	// TableToteAssignments holds tote assignments table name.
	TableToteAssignments = "tote_assignments"

	// TableConfigs holds configs table name.
	TableConfigs = "configs"

	// QueryInsertToteAssignmentLog contains query to insert ToteAssignment log.
	QueryInsertToteAssignmentLog = `
		INSERT INTO tote_assignments_log (id, state, input_data, output_data, created_at)
		VALUES (@id, @state, @input_data, @output_data, @created_at)
	`

	// QuerySelectToteAssignments contains query to select ToteAssignments.
	QuerySelectToteAssignments = `
		SELECT id, client_id, mfc_id, order_id, is_express, tote_id, lane_idx, created_at
		FROM tote_assignments
		WHERE client_id = @client_id AND mfc_id = @mfc_id ${where_part}
	`

	// QuerySelectTotesGroupedByOrder contains query to select totes grouped by order.
	QuerySelectTotesGroupedByOrder = `
		SELECT order_id, STRING_AGG(tote_id) tote_ids
		FROM tote_assignments
		WHERE client_id = @client_id AND mfc_id = @mfc_id ${where_part}
		GROUP BY order_id
	`

	// QuerySelectTotesGroupedByLane contains query to select totes grouped by lane.
	QuerySelectTotesGroupedByLane = `
		SELECT lane_idx, is_express, STRING_AGG(tote_id) tote_ids
		FROM tote_assignments
		WHERE client_id = @client_id AND mfc_id = @mfc_id ${where_part}
		GROUP BY lane_idx, is_express
	`

	// QueryInsertToteAssignment contains query to insert ToteAssignment.
	QueryInsertToteAssignment = `
		INSERT INTO ${table} (id, client_id, mfc_id, order_id, is_express, tote_id, created_at, lane_idx)
		VALUES (@id, @client_id, @mfc_id, @order_id, @is_express, @tote_id, @created_at, (${select_part}))
	`

	// QuerySelectLaneForToteAssignment contains query to select lane index for insert ToteAssignment.
	QuerySelectLaneForToteAssignment = `
		SELECT
			CASE
			-- Case1: with_free_space_on_order_lane.
			WHEN t.with_order AND t.with_free_space_on_order_lane THEN
			(
				SELECT ct1a.lane_idx FROM (
					SELECT
					tt1a.lane_idx, STRING_AGG(tt1a.order_id) order_ids, COUNT(tt1a.tote_id) len
					FROM ${table} tt1a
					WHERE
						tt1a.client_id = @client_id AND tt1a.mfc_id = @mfc_id AND tt1a.is_express = @is_express
						AND tt1a.order_id = @order_id
					GROUP BY tt1a.lane_idx
					HAVING len < @available_lane_capacity
					ORDER BY len DESC, tt1a.lane_idx ASC
					LIMIT 1
				) ct1a
			)
			-- Case2: with_free_lane.
			WHEN t.with_order AND NOT(t.with_free_space_on_order_lane) AND t.with_free_lane THEN
			(
				SELECT MAX(tt2.lane_idx)+1 FROM ${table} tt2
				WHERE tt2.client_id = @client_id AND tt2.mfc_id = @mfc_id AND tt2.is_express = @is_express
				LIMIT 1
			)
			-- Case3: not with_free_lane but with_free_space.
			-- Pay attention that with_order + with_free_space != with_free_space_on_order_lane.
			WHEN t.with_order AND NOT(t.with_free_lane) AND t.with_free_space THEN
			(
				SELECT ct1.lane_idx FROM (
					(
						SELECT
						tt1.lane_idx, STRING_AGG(tt1.order_id) order_ids, COUNT(tt1.tote_id) len
						FROM ${table} tt1
						WHERE
							tt1.client_id = @client_id AND tt1.mfc_id = @mfc_id AND tt1.is_express = @is_express
							AND tt1.order_id = @order_id
						GROUP BY tt1.lane_idx
						HAVING len < @available_lane_capacity
						ORDER BY len DESC, tt1.lane_idx ASC
						LIMIT 1
					)
					UNION ALL
					(
						SELECT
						tt1b.lane_idx, STRING_AGG(tt1b.order_id) order_ids, COUNT(tt1b.tote_id) len
						FROM ${table} tt1b
						WHERE tt1b.client_id = @client_id AND tt1b.mfc_id = @mfc_id AND tt1b.is_express = @is_express
						GROUP BY tt1b.lane_idx
						HAVING len < @available_lane_capacity
						ORDER BY len DESC, tt1b.lane_idx ASC
						LIMIT 1
					)
				) ct1 LIMIT 1
			)
			-- Case4: full ramp and with_order.
			WHEN t.with_order AND NOT(t.with_free_space) AND NOT(t.with_free_lane) THEN
			(
				SELECT ct3.lane_idx FROM (
					SELECT
					tt3.lane_idx, STRING_AGG(tt3.order_id) order_ids, COUNT(tt3.tote_id) len,
					IF(@order_id in UNNEST(ARRAY_AGG(tt3.order_id)), true, false) with_order
					FROM ${table} tt3
					WHERE tt3.client_id = @client_id AND tt3.mfc_id = @mfc_id AND tt3.is_express = @is_express
					GROUP BY tt3.lane_idx
					ORDER BY len ASC, with_order DESC, tt3.lane_idx ASC
					LIMIT 1
				) ct3
			)
			-- Case5: with_free_lane.
			WHEN NOT(t.with_order) AND t.with_free_lane THEN
			(
				-- Generate available lanes indexes.
				SELECT t4a.idx FROM (
					SELECT * FROM UNNEST(GENERATE_ARRAY(1, @available_lanes_count)) AS idx
				) t4a
				-- Select not used lane index.
				WHERE t4a.idx NOT IN (
					SELECT DISTINCT t4b.lane_idx idx FROM ${table} t4b
					WHERE t4b.client_id = @client_id AND t4b.mfc_id = @mfc_id AND t4b.is_express = @is_express
				)
				ORDER BY t4a.idx ASC
				LIMIT 1
			)
			-- Case6: no with_free_lane but with_free_space.
			WHEN NOT(t.with_order) AND NOT(t.with_free_lane) AND t.with_free_space THEN
			(
				SELECT ct5.lane_idx FROM (
					SELECT
					tt5.lane_idx, STRING_AGG(tt5.order_id) order_ids, COUNT(tt5.tote_id) len
					FROM ${table} tt5
					WHERE tt5.client_id = @client_id AND tt5.mfc_id = @mfc_id AND tt5.is_express = @is_express
					GROUP BY tt5.lane_idx
					HAVING len < @available_lane_capacity
					ORDER BY len ASC, tt5.lane_idx ASC
					LIMIT 1
				) ct5
			)
			-- Case7: full ramp and new order.
			WHEN NOT(t.with_order) AND NOT(t.with_free_lane) AND NOT(t.with_free_space) THEN
			(
				SELECT ct6.lane_idx FROM (
					SELECT
					tt6.lane_idx, STRING_AGG(tt6.order_id) order_ids, COUNT(tt6.tote_id) len
					FROM ${table} tt6
					WHERE tt6.client_id = @client_id AND tt6.mfc_id = @mfc_id AND tt6.is_express = @is_express
					GROUP BY tt6.lane_idx
					ORDER BY len ASC, tt6.lane_idx ASC
					LIMIT 1
				) ct6
			)
			-- Case8: exception, intentionally here, so we won't miss this case. 
			ELSE null
			END lane_idx
		FROM (
			SELECT
			IF((
				SELECT t1.order_id FROM ${table} t1
				WHERE
					t1.client_id = @client_id AND t1.mfc_id = @mfc_id AND t1.is_express = @is_express
					AND t1.order_id = @order_id
				LIMIT 1
			) IS null, false, true) with_order,
			IF((
				SELECT COUNT(distinct t2.lane_idx) FROM ${table} t2
				WHERE t2.client_id = @client_id AND t2.mfc_id = @mfc_id AND t2.is_express = @is_express
			) < @available_lanes_count, true, false) with_free_lane,
			IF((
				SELECT IF(COUNT(*) < @available_lane_capacity, 1, 0) wfs
				FROM ${table} t3
				WHERE t3.client_id = @client_id AND t3.mfc_id = @mfc_id AND t3.is_express = @is_express
				GROUP BY t3.lane_idx
				HAVING wfs > 0
				LIMIT 1
			) IS null, false, true) with_free_space,
			IF((
				SELECT COUNT(*)
				FROM ${table} t4
				WHERE t4.client_id = @client_id AND t4.mfc_id = @mfc_id AND t4.is_express = @is_express
				GROUP BY t4.lane_idx
				HAVING COUNT(t4.tote_id) < @available_lane_capacity AND @order_id IN UNNEST(ARRAY_AGG(t4.order_id))
				LIMIT 1
			) IS null, false, true) with_free_space_on_order_lane
		) t
	`

	// QuerySelectMFCConfigs contains query to select all MFCConfigs.
	QuerySelectMFCConfigs = `
		SELECT
			client_id, env, mfc_id, updated_at,
			error_ramp, count, depth, start, id_gen,
			lane_mapping, express_lane_mapping, flow_racks_mapping
		FROM configs
		WHERE 1=1 
	`

	// QueryUpdateLaneByToteID contains query to update lane by tote ID.
	QueryUpdateLaneByToteID = `
		UPDATE tote_assignments
		SET lane_idx = @lane_idx
		WHERE client_id = @client_id AND mfc_id = @mfc_id AND tote_id = @tote_id
	`

	// QueryDeleteToteAssignmentsByToteIDs contains query to delete ToteAssignments by tote IDs.
	QueryDeleteToteAssignmentsByToteIDs = `
		DELETE FROM tote_assignments
		WHERE client_id = @client_id AND mfc_id = @mfc_id AND tote_id IN UNNEST(@tote_ids)
	`
)

// SpannerStorage represent GCP Spanner storage service.
type SpannerStorage struct {
	client *spanner.Client
}

// NewSpannerStorage creates new GCP Spanner instance.
func NewSpannerStorage(database string) (*SpannerStorage, error) {
	c, err := spanner.NewClient(context.Background(), database)
	if err != nil {
		return nil, fmt.Errorf("failed to create new spanner client, err: %w", err)
	}

	s := &SpannerStorage{client: c}

	return s, nil
}

// GetToteAssignment {@inheritdoc}.
func (s *SpannerStorage) GetToteAssignment(
	ctx context.Context, id string,
) (out payload.ToteAssignment, err error) {
	row, err := s.client.Single().ReadRow(
		ctx, TableToteAssignments, spanner.Key{id},
		[]string{"id", "client_id", "mfc_id", "order_id", "is_express", "tote_id", "lane_idx", "created_at"},
	)
	if err != nil {
		return out, fmt.Errorf("failed to read Row, err: %w", err)
	}

	err = row.ToStruct(&out)
	if err != nil {
		return out, fmt.Errorf("failed convert row to struct, err: %w", err)
	}

	return out, nil
}

func getListToteAssignmentsStatement(input payload.ListToteAssignmentsInput, query string) spanner.Statement {
	where := ""
	if input.OrderID != "" {
		where += " AND order_id = @order_id"
	}
	query = strings.Replace(query, "${where_part}", where, -1)

	params := map[string]interface{}{
		"client_id": input.ClientID,
		"mfc_id":    input.MfcID,
		"order_id":  input.OrderID,
	}

	return spanner.Statement{SQL: query, Params: params}
}

// ListToteAssignments {@inheritdoc}.
func (s *SpannerStorage) ListToteAssignments(
	ctx context.Context, input payload.ListToteAssignmentsInput,
) ([]payload.ToteAssignment, error) {
	data := make([]payload.ToteAssignment, 0)

	stmt := getListToteAssignmentsStatement(input, QuerySelectToteAssignments)
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		r := payload.ToteAssignment{}
		err = row.Columns(
			&r.ID, &r.ClientID, &r.MfcID, &r.OrderID, &r.IsExpress, &r.ToteID, &r.LaneIdx, &r.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		data = append(data, r)
	}

	return data, nil
}

// ListTotesGroupedByOrder {@inheritdoc}.
func (s *SpannerStorage) ListTotesGroupedByOrder(
	ctx context.Context, input payload.ListToteAssignmentsInput,
) (map[string][]string, error) {
	data := make(map[string][]string)

	stmt := getListToteAssignmentsStatement(input, QuerySelectTotesGroupedByOrder)
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		orderID := ""
		totes := ""
		err = row.Columns(&orderID, &totes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		data[orderID] = strings.Split(totes, ",")
	}

	return data, nil
}

// ListTotesGroupedByLane {@inheritdoc}.
func (s *SpannerStorage) ListTotesGroupedByLane(
	ctx context.Context, input payload.ListToteAssignmentsInput,
) (out payload.ListTotesGroupedByLaneOutput, err error) {
	out.RegularLanes = make(map[int64][]string)
	out.ExpressLanes = make(map[int64][]string)

	stmt := getListToteAssignmentsStatement(input, QuerySelectTotesGroupedByLane)
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return out, fmt.Errorf("failed to get next row, err: %w", err)
		}

		var laneIdx int64
		var isExpress bool
		var totes string
		err = row.Columns(&laneIdx, &isExpress, &totes)
		if err != nil {
			return out, fmt.Errorf("failed to parse row, err: %w", err)
		}

		if isExpress {
			out.ExpressLanes[laneIdx] = strings.Split(totes, ",")
		} else {
			out.RegularLanes[laneIdx] = strings.Split(totes, ",")
		}
	}

	return out, nil
}

// LogToteAssignment {@inheritdoc}.
func (s *SpannerStorage) LogToteAssignment(
	ctx context.Context, state []payload.ToteAssignment,
	input payload.CreateToteAssignmentInput, out payload.ToteAssignment,
) error {
	id := uuid.New().String()
	params := map[string]interface{}{
		"id":          id,
		"state":       spanner.NullJSON{Value: state, Valid: true},
		"input_data":  spanner.NullJSON{Value: input, Valid: true},
		"output_data": spanner.NullJSON{Value: out, Valid: true},
		"created_at":  time.Now().UnixMilli(),
	}
	err := s.update(ctx, QueryInsertToteAssignmentLog, params)
	if err != nil {
		return fmt.Errorf("failed to perform update, err: %w", err)
	}

	return err
}

// CreateToteAssignment {@inheritdoc}.
func (s *SpannerStorage) CreateToteAssignment(
	ctx context.Context, mfcConf payload.MFCConfig, input payload.CreateToteAssignmentInput,
) (out payload.ToteAssignment, err error) {
	query := strings.Replace(
		QueryInsertToteAssignment, "${select_part}", QuerySelectLaneForToteAssignment, -1,
	)
	query = strings.Replace(query, "${table}", TableToteAssignments, -1)
	lanesCount := mfcConf.GetCount()
	if input.IsExpress {
		lanesCount = mfcConf.GetExpressCount()
	}
	params := map[string]interface{}{
		"id":                      input.ID,
		"client_id":               input.ClientID,
		"mfc_id":                  input.MfcID,
		"order_id":                input.OrderID,
		"tote_id":                 input.ToteID,
		"is_express":              input.IsExpress,
		"created_at":              time.Now().UnixMilli(),
		"available_lanes_count":   lanesCount,
		"available_lane_capacity": mfcConf.Depth,
	}
	err = s.update(ctx, query, params)
	if err != nil {
		return out, fmt.Errorf("failed to perform update, err: %w", err)
	}

	out, err = s.GetToteAssignment(ctx, input.ID)
	if err != nil {
		return out, fmt.Errorf("failed to get ToteAssignment, err: %w", err)
	}

	return out, nil
}

// GetLaneIdxForNewToteAssignment {@inheritdoc}.
func (s *SpannerStorage) GetLaneIdxForNewToteAssignment(
	ctx context.Context, mfcConf payload.MFCConfig, input payload.CreateToteAssignmentInput,
) (idx int64, err error) {
	lanesCount := mfcConf.GetCount()
	if input.IsExpress {
		lanesCount = mfcConf.GetExpressCount()
	}
	params := map[string]interface{}{
		"client_id":               input.ClientID,
		"mfc_id":                  input.MfcID,
		"order_id":                input.OrderID,
		"tote_id":                 input.ToteID,
		"is_express":              input.IsExpress,
		"available_lanes_count":   lanesCount,
		"available_lane_capacity": mfcConf.Depth,
	}

	query := strings.Replace(QuerySelectLaneForToteAssignment, "${table}", TableToteAssignments, -1)
	stmt := spanner.Statement{SQL: query, Params: params}
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to get next row, err: %w", err)
		}

		err = row.ColumnByName("lane_idx", &idx)
		if err != nil {
			return 0, fmt.Errorf("failed to parse column, err: %w", err)
		}
	}

	return idx, nil
}

// ListMFCConfigs {@inheritdoc}.
func (s *SpannerStorage) ListMFCConfigs(
	ctx context.Context, input payload.GetConfigsInput,
) ([]payload.MFCConfig, error) {
	result := make([]payload.MFCConfig, 0)

	where := ""
	if input.ClientID != "" {
		where += " AND client_id = @client_id"
	}
	if input.Env != "" {
		where += " AND env = @env"
	}
	if input.MfcID != "" {
		where += " AND mfc_id = @mfc_id"
	}
	params := map[string]interface{}{
		"client_id": input.ClientID,
		"env":       input.Env,
		"mfc_id":    input.MfcID,
	}

	stmt := spanner.Statement{SQL: QuerySelectMFCConfigs + where, Params: params}
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		r := payload.MFCConfig{}
		laneMapping := spanner.NullJSON{}
		expressLaneMapping := spanner.NullJSON{}
		flowRacksMapping := spanner.NullJSON{}
		err = row.Columns(
			&r.ClientID, &r.Env, &r.MfcID, &r.UpdatedAt,
			&r.ErrorRamp, &r.Count, &r.Depth, &r.Start, &r.IDGen,
			&laneMapping, &expressLaneMapping, &flowRacksMapping,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		rawLaneMapping := map[string]int64{}
		err = mapstructure.Decode(laneMapping.Value, &rawLaneMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal laneMapping, err: %w", err)
		}
		r.LaneMapping, err = util.ConvertKeysToInt64(rawLaneMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to convert laneMapping, err: %w", err)
		}

		rawExpressLaneMapping := map[string]int64{}
		err = mapstructure.Decode(expressLaneMapping.Value, &rawExpressLaneMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal expressLaneMapping, err: %w", err)
		}
		r.ExpressLaneMapping, err = util.ConvertKeysToInt64(rawExpressLaneMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to convert expressLaneMapping, err: %w", err)
		}

		rawFlowRacksMapping := map[string]string{}
		err = mapstructure.Decode(flowRacksMapping.Value, &rawFlowRacksMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal flowRacksMapping, err: %w", err)
		}
		r.FlowRacksMapping, err = util.ConvertKeysToInt64(rawFlowRacksMapping)
		if err != nil {
			return nil, fmt.Errorf("failed to convert flowRacksMapping, err: %w", err)
		}

		result = append(result, r)
	}

	return result, nil
}

// UpdateMFCConfig {@inheritdoc}.
func (s *SpannerStorage) UpdateMFCConfig(ctx context.Context, input payload.MFCConfig) error {
	cols := []string{
		"client_id", "env", "mfc_id", "updated_at",
		"error_ramp", "count", "depth", "start", "id_gen",
		"lane_mapping", "express_lane_mapping", "flow_racks_mapping",
	}
	vals := []interface{}{
		input.ClientID,
		input.Env,
		input.MfcID,
		input.UpdatedAt,
		input.ErrorRamp,
		input.Count,
		input.Depth,
		input.Start,
		input.IDGen,
		spanner.NullJSON{Value: input.LaneMapping, Valid: true},
		spanner.NullJSON{Value: input.ExpressLaneMapping, Valid: true},
		spanner.NullJSON{Value: input.FlowRacksMapping, Valid: true},
	}
	m := []*spanner.Mutation{spanner.InsertOrUpdate(TableConfigs, cols, vals)}
	_, err := s.client.Apply(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to upsert MFCConfig, err: %w", err)
	}

	return nil
}

// DeleteToteAssignment deletes ToteAssignment (performs delete from table).
func (s *SpannerStorage) DeleteToteAssignment(
	ctx context.Context, input payload.DeleteToteAssignmentInput,
) (out payload.DeleteToteAssignmentOutput, err error) {
	params := map[string]interface{}{
		"client_id": input.ClientID,
		"mfc_id":    input.MfcID,
		"tote_ids":  input.ToteIDs,
	}
	cb := func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{SQL: QueryDeleteToteAssignmentsByToteIDs, Params: params}
		_, err := txn.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to perform update, err: %w", err)
		}
		return nil
	}
	_, err = s.client.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return out, fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return out, nil
}

// UpdateLaneByToteID {@inheritdoc}.
func (s *SpannerStorage) UpdateLaneByToteID(ctx context.Context, input payload.UpdateToteAssignmentInput) error {
	params := map[string]interface{}{
		"client_id": input.ClientID,
		"mfc_id":    input.MfcID,
		"tote_id":   input.ToteID,
		"lane_idx":  input.LaneIdx,
	}
	err := s.update(ctx, QueryUpdateLaneByToteID, params)
	if err != nil {
		return fmt.Errorf("failed to perform update, err: %w", err)
	}

	return err
}

// update represents basic wrapper for Spanner's update method.
func (s *SpannerStorage) update(ctx context.Context, sql string, params map[string]interface{}) error {
	cb := func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{SQL: sql, Params: params}
		rowCount, err := txn.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to exec query, err: %w", err)
		}
		if rowCount < 1 {
			return fmt.Errorf("got unexpected result, rowCount: %d", rowCount)
		}
		return nil
	}
	_, err := s.client.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}
