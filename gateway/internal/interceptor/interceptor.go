package interceptor

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/missingstudio/studio/backend/internal/ratelimiter"
	"github.com/missingstudio/studio/backend/internal/schema"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
	"github.com/missingstudio/studio/common/resilience/retry"
)

var (
	ErrProviderHeaderNotExit = errors.New(fmt.Errorf("x-ms-provider provider header not available"))
	ErrRateLimitExceeded     = errors.NewForbidden("rate limit exceeded")
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

func RateLimiterInterceptor(rl *ratelimiter.RateLimiter) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			key := "req_count"
			if !rl.Limiter.Validate(key) {
				return nil, ErrRateLimitExceeded
			}

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
			provider := req.Header().Get("x-ms-provider")
			if provider == "" {
				return nil, ErrProviderHeaderNotExit
			}

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

			data := &schema.GatewayConfigHeaders{}
			err = utils.UnmarshalConfigHeaders(req.Header(), data)
			if err != nil {
				return nil, err
			}

			runner := retry.New(retry.Config{
				Times: int(data.Retry.Times),
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
