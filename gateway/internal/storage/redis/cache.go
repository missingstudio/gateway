package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheStore struct {
	client *redis.Client
	logger *slog.Logger
	wt     time.Duration
	rt     time.Duration
}

func NewCacheStore(opts ...Option) *CacheStore {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	return &CacheStore{
		client: options.client,
		logger: options.logger,
		rt:     options.rt,
		wt:     options.wt,
	}
}

func (c *CacheStore) Set(key string, value interface{}, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.wt)
	defer cancel()

	err := c.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheStore) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.rt)
	defer cancel()

	result := c.client.Get(ctx, key)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	return result.Bytes()
}
