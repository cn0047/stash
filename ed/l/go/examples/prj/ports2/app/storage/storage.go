package storage

import (
	"context"
)

// Storage describes storage service interface.
type Storage interface {
	// Create creates new key-value pair in storage.
	Create(ctx context.Context, key string, value string) error

	// Update updates existing key-value pair in storage.
	Update(ctx context.Context, key string, value string) error
}
