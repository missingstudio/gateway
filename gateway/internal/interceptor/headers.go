package interceptor

import (
	"context"
	"encoding/json"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
	"github.com/missingstudio/studio/backend/pkg/httputil"
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

			var gc httputil.GatewayConfig
			config := req.Header().Get(constants.XMSConfig)
			provider := req.Header().Get(constants.XMSProvider)

			if config != "" {
				err := json.Unmarshal([]byte(config), &gc)
				if err != nil {
					return nil, errors.ErrGatewayConfigNotValid
				}
			}

			ctx = httputil.SetContextWithProviderConfig(ctx, provider)
			ctx = httputil.SetContextWithHeaderConfig(ctx, headerConfig)
			ctx = httputil.SetContextWithGatewayConfig(ctx, &gc)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
