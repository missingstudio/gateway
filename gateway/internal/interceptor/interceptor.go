package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/config"
)

func NewLogInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func ProviderInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			provider := req.Header().Get(config.XMSProvider)
			if provider == "" {
				return nil, errors.New("x-ms-provider provider header is required")
			}

			ctx = context.WithValue(ctx, config.ProviderKey{}, provider)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
