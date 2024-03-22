package router

import (
	"log"
	"sync/atomic"
	"gateway/internal/router" // Importing to use HealthChecker
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

	// Iterate through providers to find a healthy one
	var healthyProvider *RouterConfig
	originalIdx := r.idx.Load()
	for i := 0; i < providerLen; i++ {
		idx := (originalIdx + uint64(i)) % uint64(providerLen)
		if router.DefaultHealthChecker{}.IsHealthy(r.providers[idx].Name) {
			healthyProvider = &r.providers[idx]
			r.idx.Add(1)
			break
		}
	}

	if healthyProvider == nil {
		log.Println("Error: No healthy providers available.")
		return nil
	}

	return healthyProvider
}
