package ratelimiter

import (
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowRateLimiter struct {
	logger           *slog.Logger
	cache            SlidingWindowCache
	bucketSize       int
	durationInSecond int
}

func NewSlidingWindowRateLimiter(cfg Config, logger *slog.Logger, rdb *redis.Client) RateLimiterProvider {
	return &SlidingWindowRateLimiter{
		logger:           logger,
		cache:            NewRedisLogCache(logger, rdb),
		bucketSize:       cfg.NumberOfRequests,
		durationInSecond: cfg.DurationInSecond,
	}
}

func (t *SlidingWindowRateLimiter) Validate(key string) bool {
	now := time.Now()
	window_start := now.Add(-time.Duration(t.durationInSecond) * time.Second)
	return t.cache.Validate(key, now.UnixNano(), window_start.UnixNano(), t.bucketSize)
}
