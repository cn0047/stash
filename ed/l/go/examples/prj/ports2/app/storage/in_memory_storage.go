package storage

import (
	"context"
)

// InMemoryStorage represents simple in-memory storage.
type InMemoryStorage struct {
	data map[string]string
}

// NewInMemoryStorage returns new in-memory storage instance.
func NewInMemoryStorage() *InMemoryStorage {
	i := &InMemoryStorage{}
	i.data = make(map[string]string)

	return i
}

// Create {@inheritDoc}
func (i InMemoryStorage) Create(ctx context.Context, key string, value string) error {
	// @TODO: Implement this method.
	return nil
}

// Update {@inheritDoc}
func (i InMemoryStorage) Update(ctx context.Context, key string, value string) error {
	// @TODO: Implement this method.
	return nil
}
