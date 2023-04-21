package toteassignment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/to-com/poc-td/app/mfcconfig"
	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/app/storage"
)

// ToteAssignment describes ToteAssignment service.
type ToteAssignment interface {
	// List gets ToteAssignments list.
	List(ctx context.Context, input payload.ListToteAssignmentsInput) (interface{}, error)

	// Create creates new ToteAssignment.
	Create(ctx context.Context, input payload.CreateToteAssignmentInput) (payload.CreateToteAssignmentOutput, error)

	// Delete deletes ToteAssignment.
	Delete(ctx context.Context, input payload.DeleteToteAssignmentInput) (payload.DeleteToteAssignmentOutput, error)

	// Update updates new ToteAssignment.
	Update(
		ctx context.Context, input payload.UpdateToteAssignmentInput,
	) (out payload.UpdateToteAssignmentOutput, err error)
}

// Service represents ToteAssignment service.
type Service struct {
	storage   storage.Storage
	mfcConfig mfcconfig.MFCConfig
	log       *zap.SugaredLogger
	withDBLog bool
}

// NewService creates new ToteAssignment service instance.
func NewService(
	s storage.Storage, mfcConfig mfcconfig.MFCConfig, log *zap.SugaredLogger, withDBLog bool,
) (*Service, error) {
	svc := &Service{
		storage:   s,
		log:       log,
		mfcConfig: mfcConfig,
		withDBLog: withDBLog,
	}

	return svc, nil
}

// Create {@inheritdoc}.
func (s *Service) Create(
	ctx context.Context, input payload.CreateToteAssignmentInput,
) (out payload.CreateToteAssignmentOutput, err error) {
	mfcConf, err := s.mfcConfig.Get(ctx)
	if err != nil {
		return out, fmt.Errorf("failed to get MFCConfig, err: %w", err)
	}

	if input.IsExpress && len(mfcConf.ExpressLaneMapping) == 0 {
		return out, fmt.Errorf("express lanes not enabled in MFC config")
	}

	if input.DryRun {
		idx, err := s.storage.GetLaneIdxForNewToteAssignment(ctx, mfcConf, input)
		if err != nil {
			return out, fmt.Errorf("failed to select lane index, err: %w", err)
		}
		laneID, err := s.convertLaneIdxToID(ctx, input.ClientID, input.MfcID, idx, input.IsExpress)
		if err != nil {
			return out, fmt.Errorf("failed to perform convertLaneIdxToID, err: %w", err)
		}
		out.ToteAssignment.LaneIdx = idx
		out.ToteAssignment.LaneID = laneID
		return out, nil
	}

	var state []payload.ToteAssignment
	if s.withDBLog {
		assignments, err := s.storage.ListToteAssignments(ctx, payload.ListToteAssignmentsInput{})
		if err != nil {
			s.log.Errorf("failed to perform storage.ListToteAssignments, err: %v", err)
		}
		state = assignments
	}

	if input.ID == "" {
		input.ID = uuid.New().String()
	}
	assignment, err := s.storage.CreateToteAssignment(ctx, mfcConf, input)
	if err != nil {
		return out, fmt.Errorf("failed to perform storage.CreateToteAssignment, err: %w", err)
	}
	assignment.LaneID, err = s.convertLaneIdxToID(
		ctx, assignment.ClientID, assignment.MfcID, assignment.LaneIdx, assignment.IsExpress,
	)
	if err != nil {
		return out, fmt.Errorf("failed to perform convertLaneIdxToID, err: %w", err)
	}
	out.ToteAssignment = assignment

	if s.withDBLog {
		err := s.storage.LogToteAssignment(ctx, state, input, assignment)
		if err != nil {
			s.log.Errorf("failed to perform storage.LogToteAssignments, err: %v", err)
		}
	}

	return out, nil
}

func (s *Service) convertLaneIdxToID(
	ctx context.Context, clientID string, mfcID string, idx int64, isExpress bool,
) (string, error) {
	mfcConf, err := s.mfcConfig.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get MFCConfig, err: %w", err)
	}

	id, err := mfcConf.ConvertLaneIdxToID(idx, isExpress)
	if err != nil {
		return "", fmt.Errorf("failed to ConvertLaneIdxToID: %v, %v, err: %w", clientID, mfcID, err)
	}

	return id, err
}

// List {@inheritdoc}.
func (s *Service) List(ctx context.Context, input payload.ListToteAssignmentsInput) (out interface{}, err error) {
	if input.View == payload.ViewGroupedByOrder {
		out, err = s.listToteAssignments(ctx, input)
	} else {
		out, err = s.listToteAssignmentsGroupedByOrder(ctx, input)
	}
	if err != nil {
		return out, fmt.Errorf("failed to list ToteAssignments, err: %w", err)
	}

	return out, nil
}

func (s *Service) listToteAssignments(
	ctx context.Context, input payload.ListToteAssignmentsInput,
) (out payload.ListToteAssignmentsOutput, err error) {
	out.ToteAssignments, err = s.storage.ListToteAssignments(ctx, input)
	if err != nil {
		s.log.Errorf("failed to perform storage.ListToteAssignments, err: %v", err)
	}

	for i, a := range out.ToteAssignments {
		out.ToteAssignments[i].LaneID, err = s.convertLaneIdxToID(
			ctx, a.ClientID, a.MfcID, a.LaneIdx, a.IsExpress,
		)
		if err != nil {
			return out, fmt.Errorf("failed to perform convertLaneIdxToID, err: %w", err)
		}
	}

	return out, nil
}

func (s *Service) listToteAssignmentsGroupedByOrder(
	ctx context.Context, input payload.ListToteAssignmentsInput,
) (out payload.ListToteAssignmentsGroupedByOrderOutput, err error) {
	out.Orders, err = s.storage.ListTotesGroupedByOrder(ctx, input)
	if err != nil {
		return out, fmt.Errorf("failed to perform ListTotesGroupedByOrder, err: %w", err)
	}

	lanes, err := s.storage.ListTotesGroupedByLane(ctx, input)
	if err != nil {
		return out, fmt.Errorf("failed to perform ListTotesGroupedByLane, err: %w", err)
	}

	out.Balance = make(map[string][]string, len(lanes.RegularLanes)+len(lanes.ExpressLanes))
	for laneIdx, totes := range lanes.RegularLanes {
		laneID, err := s.convertLaneIdxToID(ctx, input.ClientID, input.MfcID, laneIdx, false)
		if err != nil {
			return out, fmt.Errorf("failed to perform convertLaneIdxToID, err: %w", err)
		}
		out.Balance[laneID] = totes
	}
	for laneIdx, totes := range lanes.ExpressLanes {
		laneID, err := s.convertLaneIdxToID(ctx, input.ClientID, input.MfcID, laneIdx, true)
		if err != nil {
			return out, fmt.Errorf("failed to perform convertLaneIdxToID, err: %w", err)
		}
		out.Balance[laneID] = totes
	}

	return out, nil
}

// Delete deletes ToteAssignment.
func (s *Service) Delete(
	ctx context.Context, input payload.DeleteToteAssignmentInput) (out payload.DeleteToteAssignmentOutput, err error) {
	if len(input.ToteIDs) == 0 {
		return out, nil
	}

	out, err = s.storage.DeleteToteAssignment(ctx, input)
	if err != nil {
		return out, fmt.Errorf("failed to perform storage.DeleteToteAssignment, err: %v", err)
	}
	return out, nil
}

// Update {@inheritdoc}.
func (s *Service) Update(
	ctx context.Context, input payload.UpdateToteAssignmentInput,
) (out payload.UpdateToteAssignmentOutput, err error) {
	if input.LaneID != "" && input.ToteID != "" {
		out, err = s.updateLaneByToteID(ctx, input)
		if err != nil {
			return out, fmt.Errorf("failed to perform updateLaneByToteId, err: %w", err)
		}
	}

	return out, nil
}

func (s *Service) updateLaneByToteID(
	ctx context.Context, input payload.UpdateToteAssignmentInput,
) (out payload.UpdateToteAssignmentOutput, err error) {
	mfcConf, err := s.mfcConfig.Get(ctx)
	if err != nil {
		return out, fmt.Errorf("failed to get MFCConfig, err: %w", err)
	}

	errorRampID := mfcConf.GetErrorRampStringID()
	targetLaneID := input.LaneID
	if targetLaneID == "error" {
		targetLaneID = errorRampID
	}
	if targetLaneID != errorRampID {
		return out, fmt.Errorf("lane can be updated only to error ramp")
	}

	input.LaneIdx = int64(mfcConf.ErrorRamp)
	err = s.storage.UpdateLaneByToteID(ctx, input)
	if err != nil {
		return out, fmt.Errorf("failed to perform UpdateLaneByToteID, err: %w", err)
	}

	out.ToteAssignment.LaneIdx = input.LaneIdx
	out.ToteAssignment.LaneID = targetLaneID

	return out, nil
}
