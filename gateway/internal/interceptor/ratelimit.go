package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/internal/errors"
	"github.com/missingstudio/ai/gateway/internal/ratelimiter"
)

func RateLimiterInterceptor(rl *ratelimiter.RateLimiter) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			key := "req_count"
			if !rl.Limiter.Validate(key) {
				return nil, errors.ErrRateLimitExceeded
			}

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
