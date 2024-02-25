package ratelimiter

import (
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type IRateLimiter interface {
	Validate(key string) bool
}

type RateLimiter struct {
	Limiter IRateLimiter
}

func NewRateLimiter(rdb *redis.Client, logger *slog.Logger, rate *Rate, rltype string) *RateLimiter {
	r := &RateLimiter{}
	switch rltype {
	case "sliding_window":
		r.Limiter = NewSlidingWindowRateLimiter(rdb, logger, rate)
	default:
		r.Limiter = NewSlidingWindowRateLimiter(rdb, logger, rate)
	}

	return r
}
