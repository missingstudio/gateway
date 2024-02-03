package router

import (
	"sync/atomic"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

const (
	RoundRobin Strategy = "roundrobin"
)

type RoundRobinRouter struct {
	idx       atomic.Uint64
	providers []base.ProviderInterface
}

func NewRoundRobinRouter(providers []base.ProviderInterface) *RoundRobinRouter {
	return &RoundRobinRouter{
		providers: providers,
	}
}

func (r *RoundRobinRouter) Iterator() ProviderIterator {
	return r
}

func (r *RoundRobinRouter) Next() (base.ProviderInterface, error) {
	providerLen := len(r.providers)

	// Todo: make a check for healthy provider
	idx := r.idx.Add(1) - 1
	model := r.providers[idx%uint64(providerLen)]

	return model, nil
}
