package ratelimiter

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucketRateLimiter struct {
	logger     *slog.Logger
	cache      TokenBucketCache
	bucketSize int
	duration   time.Duration
}

func NewTokenBucketRateLimiter(cfg Config, logger *slog.Logger, rdb *redis.Client) RateLimiterProvider {
	l := &TokenBucketRateLimiter{
		logger:     logger,
		cache:      NewRedisTokenBucketCache(logger, rdb),
		bucketSize: cfg.NumberOfRequests,
		duration:   time.Duration(cfg.DurationInSecond) * time.Second,
	}
	return l
}

func (t *TokenBucketRateLimiter) Refill(key string) {
	t.cache.Refill(key, t.bucketSize)
}

func (t *TokenBucketRateLimiter) SetRoutine(key string) {
	workername := fmt.Sprintf("ratelimiter_%s", key)
	go func() {
		for {
			t.logger.Info(fmt.Sprintf("refilling %s", workername))
			t.cache.Refill(key, t.bucketSize)
			time.Sleep(time.Duration(t.duration))
		}
	}()
}

func (t *TokenBucketRateLimiter) Validate(key string) bool {
	count, isCreated := t.cache.Validate(key, t.bucketSize)
	if isCreated {
		t.SetRoutine(key)
	}
	if count > 0 {
		return true
	}
	return false
}
