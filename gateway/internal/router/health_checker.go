package router

import (
	"log"
)

type HealthChecker interface {
	IsHealthy(providerName string) bool
}

type DefaultHealthChecker struct{}

func (d *DefaultHealthChecker) IsHealthy(providerName string) bool {
	// Placeholder for actual health check logic
	// Currently returns true, assuming all providers are healthy
	return true
}
