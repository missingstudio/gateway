package interceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/internal/constants"
)

var _ connect.Interceptor = &timeoutInterceptor{}

type timeoutInterceptor struct {
	timeout time.Duration
}

// NewTimeoutInterceptor returns a new timeout interceptor.
func WithTimeout(timeout time.Duration) connect.Interceptor {
	return timeoutInterceptor{timeout: timeout}
}

func (s timeoutInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		requestTimeout := s.timeout

		timeoutHeader := req.Header().Get(constants.XMSRequestTimeout)
		timeoutDuration, err := time.ParseDuration(timeoutHeader)
		// If timeout is provided in the header, use it; otherwise, use the default request timeout
		if err == nil {
			requestTimeout = timeoutDuration
		}

		ctx, cancel := context.WithTimeout(ctx, requestTimeout)
		defer cancel()

		return next(ctx, req)
	})
}

func (s timeoutInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		ctx, cancel := context.WithTimeout(ctx, s.timeout)
		defer cancel()
		return next(ctx, spec)
	})
}

func (s timeoutInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, shc connect.StreamingHandlerConn) error {
		requestTimeout := s.timeout

		timeoutHeader := shc.RequestHeader().Get(constants.XMSRequestTimeout)
		timeoutDuration, err := time.ParseDuration(timeoutHeader)
		// If timeout is provided in the header, use it; otherwise, use the default request timeout
		if err != nil {
			requestTimeout = timeoutDuration
		}

		ctx, cancel := context.WithTimeout(ctx, requestTimeout)
		defer cancel()

		return next(ctx, shc)
	})
}
