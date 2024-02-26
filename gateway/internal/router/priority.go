package router

import (
	"sync/atomic"
)

const (
	Priority Strategy = "priority"
)

type PriorityRouter struct {
	idx       *atomic.Uint64
	providers []RouterConfig
}

func NewPriorityRouter(providers []RouterConfig) *PriorityRouter {
	return &PriorityRouter{
		idx:       &atomic.Uint64{},
		providers: providers,
	}
}

func (r *PriorityRouter) Next() (*RouterConfig, error) {
	idx := int(r.idx.Load())

	// Todo: make a check for healthy provider
	model := &r.providers[idx]
	r.idx.Add(1)

	return model, nil
}
