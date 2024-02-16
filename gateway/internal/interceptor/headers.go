package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
)

func ProviderInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Check if required headers are available
			provider := req.Header().Get(constants.XMSProvider)
			config := req.Header().Get(constants.XMSConfig)
			if provider == "" && config == "" {
				return nil, errors.ErrRequiredHeaderNotExit
			}

			// Check if provider has registered of not
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
