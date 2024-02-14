package ratelimiter

import (
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	Limiter RateLimiterProvider
}

func NewRateLimiter(cfg Config, logger *slog.Logger, rltype string, rdb *redis.Client) *RateLimiter {
	r := &RateLimiter{}
	switch rltype {
	case "sliding_window":
		r.Limiter = NewSlidingWindowRateLimiter(cfg, logger, rdb)
	default:
		r.Limiter = NewSlidingWindowRateLimiter(cfg, logger, rdb)
	}

	return r
}
