package httputil

import (
	"context"
)

type (
	HeaderConfigContextKey struct{}
	headerConfigContext    map[string]any
)

func SetContextWithHeaderConfig(ctx context.Context, config headerConfigContext) context.Context {
	return context.WithValue(ctx, HeaderConfigContextKey{}, config)
}

func GetContextWithHeaderConfig(ctx context.Context) headerConfigContext {
	c, ok := ctx.Value(HeaderConfigContextKey{}).(headerConfigContext)
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
