package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/missingstudio/ai/gateway/internal/constants"
	"github.com/missingstudio/ai/gateway/internal/errors"
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

type RouterConfig struct {
	Name       string         `json:"name"`
	Auth       map[string]any `json:"auth"`
	VirtualKey string         `json:"virtual_key"`
	Retry      RetryConfig    `json:"retry"`
	Cache      CacheConfig    `json:"cache"`
	Strategy   StrategyConfig `json:"strategy"`
	Targets    []RouterConfig `json:"targets"`
	Metadata   map[string]any `json:"metadata"`
}

func NewRouterConfig(config string, headers http.Header) (*RouterConfig, error) {
	rc := &RouterConfig{}

	// Unmarshal the config string into rc if it's not empty
	if config != "" {
		if err := json.Unmarshal([]byte(config), rc); err != nil {
			return nil, errors.ErrRouterConfigNotValid
		}
	}

	// Retrieve values from headers
	provider := headers.Get(constants.XMSProvider)
	authorization := headers.Get(constants.Authorization)

	// Update router config based on headers
	if provider != "" {
		rc.Name = provider
	}

	if authorization != "" {
		// Ensure that Auth map is initialized before updating it
		if rc.Auth == nil {
			rc.Auth = make(map[string]interface{})
		}

		rc.Auth["api_key"] = authorization
	}

	return rc, nil
}
