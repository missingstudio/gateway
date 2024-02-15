package cache

import (
	"context"
	"time"
)

type ICache interface {
	// Get a key from cache.
	Get(ctx context.Context, key string) ([]byte, error)
	// Set a key in cache.
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	// Delete a key in cache.
	Delete(ctx context.Context, key string) error
	// IsAlive performs a healthcheck on the cache.
	IsAlive(context.Context) (bool, error)
}
