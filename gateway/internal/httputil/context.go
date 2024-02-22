package httputil

import (
	"context"

	"github.com/missingstudio/studio/backend/internal/schema"
)

type (
	GatewayConfigContextKey struct{}
	HeaderConfigContextKey  struct{}
)

type headerConfig map[string]any

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

func SetContextWithGatewayConfig(ctx context.Context, config *schema.GatewayConfig) context.Context {
	return context.WithValue(ctx, GatewayConfigContextKey{}, config)
}

func GetContextWithGatewayConfig(ctx context.Context) *schema.GatewayConfig {
	c, ok := ctx.Value(GatewayConfigContextKey{}).(*schema.GatewayConfig)
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
