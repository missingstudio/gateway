package api

import (
	"log/slog"

	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
)

type Deps struct {
	Logger      *slog.Logger
	Ingester    ingester.Ingester
	RateLimiter *ratelimiter.RateLimiter
}

func NewDeps(logger *slog.Logger, ingester ingester.Ingester, ratelimiter *ratelimiter.RateLimiter) Deps {
	return Deps{
		Logger:      logger,
		Ingester:    ingester,
		RateLimiter: ratelimiter,
	}
}
