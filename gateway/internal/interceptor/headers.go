package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/internal/constants"
	"github.com/missingstudio/ai/gateway/internal/errors"
	"github.com/missingstudio/ai/gateway/internal/router"
)

var _ connect.Interceptor = &headersInterceptor{}

type headersInterceptor struct{}

func WithHeaderConfig() connect.Interceptor {
	return &headersInterceptor{}
}

func (h *headersInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		config := req.Header().Get(constants.XMSConfig)

		rc, err := router.NewRouterConfig(config, req.Header())
		if err != nil {
			return nil, errors.ErrRouterConfigNotValid
		}

		ctx = router.SetContextWithRouterConfig(ctx, rc)
		return next(ctx, req)
	})
}

func (h *headersInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return next(ctx, spec)
	})
}

func (h *headersInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, shc connect.StreamingHandlerConn) error {
		config := shc.RequestHeader().Get(constants.XMSConfig)

		rc, err := router.NewRouterConfig(config, shc.RequestHeader())
		if err != nil {
			return errors.ErrRouterConfigNotValid
		}

		ctx = router.SetContextWithRouterConfig(ctx, rc)
		return next(ctx, shc)
	})
}
