package schema

import (
	"time"
)

type RetryConfig struct {
	// Number is the number of times to retry the request when a retryable
	Numbers int32 `json:"numbers" default:"1"`

	// RetryOnStatusCodes is a flat list of http response status codes that are
	// eligible for retry. This again should be feasible in any reasonable proxy.
	OnStatusCodes []uint32 `json:"on_status_codes" default:"[]"`
}

type CacheConfig struct {
	// Mode specifies the type of cache with two possible modes: simple and semantic.
	Mode string `json:"mode" default:"simple"`
	// TTL (Time To Live) is the duration for which cache entries should be considered valid.
	TTL time.Duration `json:"ttl"`
}

type StrategyConfig struct {
	Mode string `json:"mode"`
}

type GatewayConfig struct {
	Name        string          `json:"name"`
	ApiKey      string          `json:"api_key"`
	VirtualKey  string          `json:"virtual_key"`
	RetryConfig RetryConfig     `json:"retry"`
	CacheConfig CacheConfig     `json:"cache"`
	Targets     []GatewayConfig `json:"targets"`
	Metadata    map[string]any  `josn:"metadata"`
}

func DefaultGatewayConfig() *GatewayConfig {
	return &GatewayConfig{}
}
