package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/httputil"
	"github.com/missingstudio/studio/common/resilience/retry"
)

func RetryInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			var err error
			var response connect.AnyResponse
			config := httputil.GetContextWithGatewayConfig(ctx)
			runner := retry.New(retry.Config{
				Numbers: int(config.RetryConfig.Numbers),
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
