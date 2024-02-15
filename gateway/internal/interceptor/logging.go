package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

var _ connect.Interceptor = &loggingInterceptor{}

type loggingInterceptor struct {
	logger *slog.Logger
}

func NewLoggingInterceptor(logger *slog.Logger) *loggingInterceptor {
	return &loggingInterceptor{
		logger,
	}
}

func (l *loggingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		if err != nil {
			l.logger.Error("response with error",
				"err", err.Error(),
				"endpoint", req.Spec().Procedure,
				"addr", req.Peer().Addr)
		}
		return res, err
	})
}

func (i *loggingInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		conn := next(ctx, spec)
		return conn
	}
}

func (i *loggingInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		return next(ctx, conn)
	}
}
