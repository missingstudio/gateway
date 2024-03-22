package router

import (
	"sync/atomic"
	"log"
	"gateway/internal/router" // Importing to use HealthChecker
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
	providerLen := len(r.providers)
	originalIdx := r.idx.Load()
	var healthyProvider *RouterConfig
	for i := 0; i < providerLen; i++ {
		idx := (originalIdx + uint64(i)) % uint64(providerLen)
		if router.DefaultHealthChecker{}.IsHealthy(r.providers[idx].Name) {
			healthyProvider = &r.providers[idx]
			r.idx.Store(idx + 1)
			break
		}
	}
	if healthyProvider == nil {
		log.Println("Error: No healthy providers available.")
		return nil, fmt.Errorf("no healthy providers available")
	}
	return healthyProvider, nil
}
