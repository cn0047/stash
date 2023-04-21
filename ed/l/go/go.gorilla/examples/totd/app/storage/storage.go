package storage

import (
	"context"

	"github.com/to-com/poc-td/app/payload"
)

// Storage describes storage service interface.
type Storage interface {
	// ListToteAssignments gets ToteAssignments.
	ListToteAssignments(
		ctx context.Context, input payload.ListToteAssignmentsInput,
	) ([]payload.ToteAssignment, error)

	// ListTotesGroupedByOrder gets totes grouped by order.
	ListTotesGroupedByOrder(
		ctx context.Context, input payload.ListToteAssignmentsInput,
	) (map[string][]string, error)

	// ListTotesGroupedByLane gets totes grouped by lane.
	ListTotesGroupedByLane(
		ctx context.Context, input payload.ListToteAssignmentsInput,
	) (payload.ListTotesGroupedByLaneOutput, error)

	// GetToteAssignment gets ToteAssignment by ID.
	GetToteAssignment(ctx context.Context, id string) (payload.ToteAssignment, error)

	// CreateToteAssignment creates ToteAssignment (performs insert into table).
	CreateToteAssignment(
		ctx context.Context, mfcConf payload.MFCConfig, input payload.CreateToteAssignmentInput,
	) (payload.ToteAssignment, error)

	// UpdateLaneByToteID updates lane using tote ID to determine target ToteAssignment.
	UpdateLaneByToteID(ctx context.Context, input payload.UpdateToteAssignmentInput) error

	// GetLaneIdxForNewToteAssignment represents dry-run for CreateToteAssignment.
	GetLaneIdxForNewToteAssignment(
		ctx context.Context, mfcConf payload.MFCConfig, input payload.CreateToteAssignmentInput,
	) (int64, error)

	// LogToteAssignment creates record in ToteAssignments log table.
	LogToteAssignment(
		ctx context.Context, state []payload.ToteAssignment,
		input payload.CreateToteAssignmentInput, out payload.ToteAssignment,
	) error

	// DeleteToteAssignment deletes ToteAssignment.
	DeleteToteAssignment(ctx context.Context, input payload.DeleteToteAssignmentInput,
	) (payload.DeleteToteAssignmentOutput, error)

	// ListMFCConfigs gets MFCConfigs.
	ListMFCConfigs(ctx context.Context, input payload.GetConfigsInput) ([]payload.MFCConfig, error)

	// UpdateMFCConfig updates MFCConfig.
	UpdateMFCConfig(ctx context.Context, input payload.MFCConfig) error
}
