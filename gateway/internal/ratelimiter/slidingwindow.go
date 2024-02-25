package ratelimiter

import (
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowRateLimiter struct {
	logger *slog.Logger
	cache  SlidingWindowCache
	limit  int
	period time.Duration
}

func NewSlidingWindowRateLimiter(rdb *redis.Client, logger *slog.Logger, rate *Rate) IRateLimiter {
	return &SlidingWindowRateLimiter{
		logger: logger,
		cache:  NewRedisLogCache(logger, rdb),
		limit:  rate.Limit,
		period: rate.Period,
	}
}

func (t *SlidingWindowRateLimiter) Validate(key string) bool {
	now := time.Now()
	window_start := now.Add(-t.period * time.Second)
	return t.cache.Validate(key, now.UnixNano(), window_start.UnixNano(), t.limit)
}
