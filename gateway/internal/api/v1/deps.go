package v1

import (
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
)

type Deps struct {
	ingester    ingester.Ingester
	ratelimiter *ratelimiter.RateLimiter
}

func NewDeps(ingester ingester.Ingester, ratelimiter *ratelimiter.RateLimiter) *Deps {
	return &Deps{
		ingester:    ingester,
		ratelimiter: ratelimiter,
	}
}
