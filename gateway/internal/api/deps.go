package api

import (
	"log/slog"

	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	"github.com/missingstudio/ai/gateway/internal/providers"
	"github.com/missingstudio/ai/gateway/internal/ratelimiter"
)

type Deps struct {
	Logger            *slog.Logger
	Ingester          ingester.Ingester
	RateLimiter       *ratelimiter.RateLimiter
	ProviderService   *providers.Service
	ConnectionService *connection.Service
	PromptService     *prompt.Service
	APIKeyService     *apikey.Service
	AuthEnabled       bool
}

func NewDeps(
	logger *slog.Logger,
	ingester ingester.Ingester,
	ratelimiter *ratelimiter.RateLimiter,
	ps *providers.Service,
	cs *connection.Service,
	pms *prompt.Service,
	aks *apikey.Service,
	authEnabled bool,
) *Deps {
	return &Deps{
		Logger:            logger,
		Ingester:          ingester,
		RateLimiter:       ratelimiter,
		ProviderService:   ps,
		ConnectionService: cs,
		PromptService:     pms,
		APIKeyService:     aks,
		AuthEnabled:       authEnabled,
	}
}
