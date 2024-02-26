package router

import (
	"sync/atomic"
)

const (
	RoundRobin Strategy = "roundrobin"
)

type RoundRobinRouter struct {
	idx       atomic.Uint64
	providers []RouterConfig
}

func NewRoundRobinRouter(providers []RouterConfig) RouterIterator {
	return &RoundRobinRouter{
		providers: providers,
	}
}

func (r *RoundRobinRouter) Iterator() RouterIterator {
	return r
}

func (r *RoundRobinRouter) Next() *RouterConfig {
	providerLen := len(r.providers)

	// Todo: make a check for healthy provider
	idx := r.idx.Add(1) - 1
	model := &r.providers[idx%uint64(providerLen)]

	return model
}
