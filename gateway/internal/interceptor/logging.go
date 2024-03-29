package interceptor

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/missingstudio/common/logger"
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

		resultStatus := http.StatusOK
		if err != nil {
			if err, ok := err.(*connect.Error); ok {
				switch err.Code() {
				case connect.CodeInternal, connect.CodeUnknown:
					resultStatus = http.StatusInternalServerError
				default:
					resultStatus = http.StatusBadRequest
				}
			}
		}

		defer func() {
			l := logger.NewConnectRequestLogger(l.logger, resultStatus, req, res)
			l.ConnectRequestf(ctx)
		}()
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
