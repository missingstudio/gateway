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
	providers []base.IProvider
}

func NewRoundRobinRouter(providers []base.IProvider) *RoundRobinRouter {
	return &RoundRobinRouter{
		providers: providers,
	}
}

func (r *RoundRobinRouter) Iterator() ProviderIterator {
	return r
}

func (r *RoundRobinRouter) Next() (base.IProvider, error) {
	providerLen := len(r.providers)

	// Todo: make a check for healthy provider
	idx := r.idx.Add(1) - 1
	model := r.providers[idx%uint64(providerLen)]

	return model, nil
}
