package httputil

import (
	"context"
)

type (
	GatewayConfigContextKey struct{}
	HeaderConfigContextKey  struct{}
)

type headerConfig map[string]any

type RetryConfig struct {
	// Number is the number of times to retry the request when a retryable
	Number int32 `json:"number"`

	// RetryOnStatusCodes is a flat list of http response status codes that are
	// eligible for retry. This again should be feasible in any reasonable proxy.
	OnStatusCodes []uint32 `json:"on_status_codes"`
}

type CacheConfig struct {
	Mode   string `json:"mode"`
	MaxAge int32  `json:"max_age"`
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

type Address struct {
	City string `json:"city"`
	Zip  string `json:"zip"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func SetContextWithHeaderConfig(ctx context.Context, config headerConfig) context.Context {
	return context.WithValue(ctx, HeaderConfigContextKey{}, config)
}

func GetContextWithHeaderConfig(ctx context.Context) headerConfig {
	c, ok := ctx.Value(HeaderConfigContextKey{}).(headerConfig)
	if !ok {
		return nil
	}
	return c
}

func SetContextWithGatewayConfig(ctx context.Context, config *GatewayConfig) context.Context {
	return context.WithValue(ctx, GatewayConfigContextKey{}, config)
}

func GetContextWithGatewayConfig(ctx context.Context) *GatewayConfig {
	c, ok := ctx.Value(GatewayConfigContextKey{}).(*GatewayConfig)
	if !ok {
		return nil
	}
	return c
}

func SetContextWithProviderConfig(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, HeaderConfigContextKey{}, name)
}

func GetContextWithProviderConfig(ctx context.Context) string {
	name, ok := ctx.Value(HeaderConfigContextKey{}).(string)
	if !ok {
		return ""
	}
	return name
}
