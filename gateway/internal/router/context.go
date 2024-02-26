package router

import (
	"context"
)

type (
	RouterServiceContextKey struct{}
	RouterConfigContextKey  struct{}
)

func SetContextWithHeaderConfig(ctx context.Context, config *RouterService) context.Context {
	return context.WithValue(ctx, RouterServiceContextKey{}, config)
}

func GetContextWithHeaderConfig(ctx context.Context) *RouterService {
	c, ok := ctx.Value(RouterServiceContextKey{}).(*RouterService)
	if !ok {
		return nil
	}
	return c
}

func SetContextWithRouterConfig(ctx context.Context, config *RouterConfig) context.Context {
	return context.WithValue(ctx, RouterConfigContextKey{}, config)
}

func GetContextWithRouterConfig(ctx context.Context) *RouterConfig {
	c, ok := ctx.Value(RouterConfigContextKey{}).(*RouterConfig)
	if !ok {
		return nil
	}
	return c
}
