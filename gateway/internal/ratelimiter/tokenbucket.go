package ratelimiter

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucketRateLimiter struct {
	logger *slog.Logger
	cache  TokenBucketCache
	limit  int
	period time.Duration
}

func NewTokenBucketRateLimiter(rdb *redis.Client, logger *slog.Logger, rate *Rate) IRateLimiter {
	l := &TokenBucketRateLimiter{
		logger: logger,
		cache:  NewRedisTokenBucketCache(logger, rdb),
		limit:  rate.Limit,
		period: rate.Period,
	}
	return l
}

func (t *TokenBucketRateLimiter) Refill(key string) {
	t.cache.Refill(key, t.limit)
}

func (t *TokenBucketRateLimiter) SetRoutine(key string) {
	workername := fmt.Sprintf("ratelimiter_%s", key)
	go func() {
		for {
			t.logger.Info(fmt.Sprintf("refilling %s", workername))
			t.cache.Refill(key, t.limit)
			time.Sleep(t.period)
		}
	}()
}

func (t *TokenBucketRateLimiter) Validate(key string) bool {
	count, isCreated := t.cache.Validate(key, t.limit)
	if isCreated {
		t.SetRoutine(key)
	}
	if count > 0 {
		return true
	}
	return false
}
