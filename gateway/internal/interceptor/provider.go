package interceptor

import (
	"context"

	"connectrpc.com/connect"
)

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
