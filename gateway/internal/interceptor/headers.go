package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
	"github.com/missingstudio/studio/backend/internal/httputil"
	"github.com/missingstudio/studio/backend/internal/router"
)

func HeadersInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			headerConfig := make(map[string]any)
			for key, values := range req.Header() {
				if len(values) > 0 {
					headerConfig[key] = values[0]
				}
			}

			config := req.Header().Get(constants.XMSConfig)
			rc, err := router.NewRouterConfig(config)
			if err != nil {
				return nil, errors.ErrRouterConfigNotValid
			}

			provider := req.Header().Get(constants.XMSProvider)
			ctx = httputil.SetContextWithProviderConfig(ctx, provider)
			ctx = httputil.SetContextWithHeaderConfig(ctx, headerConfig)
			ctx = router.SetContextWithRouterConfig(ctx, rc)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
