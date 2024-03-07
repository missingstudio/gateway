package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/common/resilience/retry"
	"github.com/missingstudio/ai/gateway/internal/router"
)

func RetryInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			var err error
			var response connect.AnyResponse
			config := router.GetContextWithRouterConfig(ctx)
			runner := retry.New(retry.Config{
				Numbers: int(config.Retry.Numbers),
			})

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
