package ports

import (
	"context"
	"encoding/json"
	"fmt"

	"ptt/app/payload"
	"ptt/app/storage"
)

var (
	// ErrPortIDLength represents error when port ID is too short.
	ErrPortIDLength = fmt.Errorf("port ID must be at least 3 characters long")
)

// PortsService represents service to work with ports.
type PortsService struct {
	Storage storage.Storage
}

// New returns new ports service instance.
func New(storage storage.Storage) *PortsService {
	p := &PortsService{
		Storage: storage,
	}

	return p
}

// InitFromFile initializes ports service by reading JSON file and saving data into storage service.
// @TDB.
func (p *PortsService) InitFromFile(filePath string) {
	// @TODO: Implement this method.
}

// CreatePort creates new port.
func (p *PortsService) CreatePort(ctx context.Context, port *payload.Port) (*payload.Port, error) {
	// Some Business Logic here...

	// @TODO/TBD: Maybe check whether key exists in storage.

	// @TBD: Something like part of business logic, which will be used in tests later.
	if len(port.ID) < 3 {
		return nil, ErrPortIDLength
	}

	key := port.ID
	val, err := json.Marshal(p) // Just for purpose of test task KISS applied here.
	if err != nil {
		return nil, fmt.Errorf("failed to serialize port, err: %w", err)
	}
	err = p.Storage.Create(ctx, key, string(val))
	if err != nil {
		return nil, fmt.Errorf("failed to create port, err: %w", err)
	}

	// Some Business Logic here...

	return port, nil
}

// UpdatePort updates existing port.
func (p *PortsService) UpdatePort(ctx context.Context, port *payload.Port) (*payload.Port, error) {
	// Some Business Logic here...

	// @TODO/TBD: Maybe check whether key exists in storage.

	key := port.ID
	val, err := json.Marshal(p) // Just for purpose of test task KISS applied here.
	if err != nil {
		return nil, fmt.Errorf("failed to serialize port, err: %w", err)
	}
	err = p.Storage.Update(ctx, key, string(val))
	if err != nil {
		return nil, fmt.Errorf("failed to create port, err: %w", err)
	}

	// Some Business Logic here...

	return port, nil
}
