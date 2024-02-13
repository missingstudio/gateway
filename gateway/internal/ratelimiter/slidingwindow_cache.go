package ratelimiter

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowCache interface {
	Validate(keyname string, timestamp int64, window_start int64, limit int) bool
}

type redisSlidingWindowCache struct {
	logger *slog.Logger
	client *redis.Client
}

func NewRedisLogCache(logger *slog.Logger, client *redis.Client) SlidingWindowCache {
	return &redisSlidingWindowCache{
		logger: logger,
		client: client,
	}
}

func (c *redisSlidingWindowCache) Validate(keyname string, timestamp int64, window_start int64, limit int) bool {
	script := redis.NewScript(`
		local keyname = KEYS[1]
		local now = tonumber(ARGV[1])
		local window_start = tonumber(ARGV[2])
		local limit = tonumber(ARGV[3])

		redis.call('ZREMRANGEBYSCORE', keyname, 0, window_start)

		local count = redis.call('ZCARD', keyname)
		if count < limit then
		redis.call('ZADD', keyname, now, now)
		end

		return limit - count
	`)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n, err := script.Run(ctx, c.client, []string{keyname}, timestamp, window_start, limit).Result()
	if err != nil {
		c.logger.Error("failed to validate rate limiter")
		return false
	}

	if val, ok := n.(int64); ok && val > 0 {
		return true
	}
	return false
}
