package schema

import (
	"time"
)

type CacheConfig struct {
	// Mode specifies the type of cache with two possible modes: simple and semantic.
	Mode string `json:"mode" default:"simple"`
	// TTL (Time To Live) is the duration for which cache entries should be considered valid.
	TTL time.Duration `json:"ttl"`
}

type RetryConfig struct {
	Times int32 `json:"times" default:"1"`
	// Status codes for retry
	OnStatusCodes []string `json:"on_status_codes"`
}

type StrategyConfig struct {
	Mode string `json:"mode" default:"fallback"`
	// Status codes for retry
	OnStatusCodes []string `json:"on_status_codes"`
}

type GatewayConfigHeaders struct {
	// Virtual key is temporary key with configurations for the gateway
	Provider   string `json:"provider"`
	VirtualKey string `json:"virtual_key"`
	// Cache represents the cache configuration for the gateway.
	Cache     CacheConfig    `json:"cache"`
	Retry     RetryConfig    `json:"retry"`
	Strategy  StrategyConfig `json:"strategy"`
	Providers []any          `json:"providers"`
}
