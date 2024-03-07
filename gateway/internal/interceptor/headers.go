package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/internal/constants"
	"github.com/missingstudio/ai/gateway/internal/errors"
	"github.com/missingstudio/ai/gateway/internal/router"
)

func HeadersInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			config := req.Header().Get(constants.XMSConfig)

			rc, err := router.NewRouterConfig(config, req.Header())
			if err != nil {
				return nil, errors.ErrRouterConfigNotValid
			}

			ctx = router.SetContextWithRouterConfig(ctx, rc)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
