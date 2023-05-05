package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

// RedisStorage represents redis storage.
type RedisStorage struct {
	Client *redis.Client
}

// NewRedisStorage returns new redis storage instance.
func NewRedisStorage(host string, port string, password string, db int) *RedisStorage {
	r := &RedisStorage{}

	addr := fmt.Sprintf("%s:%s", host, port)
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return r
}

// Create {@inheritDoc}
func (r RedisStorage) Create(ctx context.Context, key string, value string) error {
	// @TODO/TDB: Maybe read key beforehand, and error in case key exists.

	err := r.Client.Set(key, value, 0).Err() // Just for purpose of test task KISS applied here.
	if err != nil {
		return fmt.Errorf("failed to create redis key-value pair, err: %w", err)
	}

	return nil
}

// Update {@inheritDoc}
func (r RedisStorage) Update(ctx context.Context, key string, value string) error {
	// @TODO/TDB: Maybe read key beforehand, and error in case key not exists.

	err := r.Client.Set(key, value, 0).Err() // Just for purpose of test task KISS applied here.
	if err != nil {
		return fmt.Errorf("failed to update redis key-value pair, err: %w", err)
	}

	return nil
}
