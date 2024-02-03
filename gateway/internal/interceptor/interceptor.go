package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/common/resilience/retry"
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

func RetryInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			var err error
			var response connect.AnyResponse

			runner := retry.New(retry.Config{})
			err = runner.Run(ctx, func(ctx context.Context) error {
				response, err = next(ctx, req)
				if err != nil {
					return err
				}
				return nil
			})

			return response, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
