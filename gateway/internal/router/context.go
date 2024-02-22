package router

import (
	"context"
)

type (
	RouterServiceContextKey struct{}
	RouterConfigContextKey  struct{}
)

type (
	RouterServiceContext *RouterService
	routerConfigContext  *RouterConfig
)

func SetContextWithHeaderConfig(ctx context.Context, config RouterServiceContext) context.Context {
	return context.WithValue(ctx, RouterServiceContextKey{}, config)
}

func GetContextWithHeaderConfig(ctx context.Context) RouterServiceContext {
	c, ok := ctx.Value(RouterServiceContextKey{}).(RouterServiceContext)
	if !ok {
		return nil
	}
	return c
}

func SetContextWithRouterConfig(ctx context.Context, config routerConfigContext) context.Context {
	return context.WithValue(ctx, RouterConfigContextKey{}, config)
}

func GetContextWithRouterConfig(ctx context.Context) routerConfigContext {
	c, ok := ctx.Value(RouterConfigContextKey{}).(routerConfigContext)
	if !ok {
		return nil
	}
	return c
}
