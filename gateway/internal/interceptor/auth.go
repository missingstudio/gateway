package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
)

// NewAPIKeyInterceptor returns interceptor which is checking if api key exits
func NewAPIKeyInterceptor(logger *slog.Logger) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			apiHeader := req.Header().Get(constants.XMSAPIKey)
			if apiHeader == "" {
				logger.Info("request without api key",
					"api_key", apiHeader,
					"addr", req.Peer().Addr,
					"endpoint", req.Spec().Procedure)
				return nil, errors.ErrUnauthenticated
			}

			return next(ctx, req)
		})
	})
}
