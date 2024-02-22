package router

import (
	"encoding/json"
	"time"

	"github.com/missingstudio/studio/backend/internal/errors"
	"github.com/missingstudio/studio/backend/internal/providers/base"
)

type (
	Strategy string
	IRouter  interface{}
)

type RouterIterator interface {
	Next() (base.IProvider, error)
}

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
	Name        string         `json:"name"`
	ApiKey      string         `json:"api_key"`
	VirtualKey  string         `json:"virtual_key"`
	RetryConfig RetryConfig    `json:"retry"`
	CacheConfig CacheConfig    `json:"cache"`
	Targets     []RouterConfig `json:"targets"`
	Metadata    map[string]any `josn:"metadata"`
}

func NewRouterConfig(config string) (*RouterConfig, error) {
	rc := &RouterConfig{}

	if config != "" {
		err := json.Unmarshal([]byte(config), rc)
		if err != nil {
			return nil, errors.ErrRouterConfigNotValid
		}
	}
	return &RouterConfig{}, nil
}
