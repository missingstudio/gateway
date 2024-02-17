package api

import (
	"log/slog"

	"github.com/missingstudio/studio/backend/internal/connections"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
)

type Deps struct {
	Logger            *slog.Logger
	Ingester          ingester.Ingester
	RateLimiter       *ratelimiter.RateLimiter
	ProviderService   *providers.Service
	ConnectionService *connections.Service
}

func NewDeps(
	logger *slog.Logger,
	ingester ingester.Ingester,
	ratelimiter *ratelimiter.RateLimiter,
	ps *providers.Service,
	cs *connections.Service,
) *Deps {
	return &Deps{
		Logger:            logger,
		Ingester:          ingester,
		RateLimiter:       ratelimiter,
		ProviderService:   ps,
		ConnectionService: cs,
	}
}
