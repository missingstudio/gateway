package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/durationpb"
)

type (
	SpanIDKey      struct{}
	TraceIDKey     struct{}
	RequestTimeKey struct{}
)

type Option func(*slog.HandlerOptions)

func WithLevel(level slog.Level) Option {
	return func(opts *slog.HandlerOptions) {
		opts.Level = level
	}
}

func New(formatAsJson bool, options ...Option) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{}

	if len(options) > 0 {
		for _, option := range options {
			option(handlerOptions)
		}
	}

	if formatAsJson {
		return slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
}

type httpRequest struct {
	RequestMethod string               `json:"requestMethod,omitempty"`
	RequestUrl    string               `json:"requestUrl,omitempty"`
	RequestSize   string               `json:"requestSize,omitempty"`
	Status        int                  `json:"status,omitempty"`
	ResponseSize  string               `json:"responseSize,omitempty"`
	UserAgent     string               `json:"userAgent,omitempty"`
	Protocol      string               `json:"protocol,omitempty"`
	RemoteIp      string               `json:"remoteIp,omitempty"`
	ServerIp      string               `json:"serverIp,omitempty"`
	Referer       string               `json:"referer,omitempty"`
	Latency       *durationpb.Duration `json:"latency,omitempty"`
}

type connectRequestLogger struct {
	Logger *slog.Logger
	Status int
	Req    connect.AnyRequest
	Resp   connect.AnyResponse
}

func NewConnectRequestLogger(l *slog.Logger, status int, req connect.AnyRequest, resp connect.AnyResponse) *connectRequestLogger {
	return &connectRequestLogger{
		Logger: l,
		Status: status,
		Req:    req,
		Resp:   resp,
	}
}

func defaultLogAttrs(severity slog.Level, traceID, spanID string) []slog.Attr {
	return []slog.Attr{
		slog.String("severity", severity.String()),
		slog.String("spanId", spanID),
		slog.String("traceId", traceID),
	}
}

func (l *connectRequestLogger) ConnectRequestf(ctx context.Context) {
	req := l.Req

	spanID, ok := ctx.Value(SpanIDKey{}).(string)
	if !ok {
		spanID = ""
	}
	traceID, ok := ctx.Value(TraceIDKey{}).(string)
	if !ok {
		traceID = ""
	}

	requestTime, ok := ctx.Value(RequestTimeKey{}).(time.Time)
	if !ok {
		requestTime = time.Now()
	}

	msg := "Connect request info"
	attrs := append(
		defaultLogAttrs(slog.LevelInfo, traceID, spanID),
		slog.Any("httpRequest", httpRequest{
			RequestMethod: req.HTTPMethod(),
			Status:        l.Status,
			RequestUrl:    "https://" + req.Header().Get("Host") + req.Spec().Procedure,
			UserAgent:     req.Header().Get("User-Agent"),
			Protocol:      req.Header().Get("Protocol"),
			RemoteIp:      req.Header().Get("X-Forwarded-For"),
			ServerIp:      req.Peer().Addr,
			Latency:       durationpb.New(time.Since(requestTime)),
		}),
		slog.Time("time", requestTime),
	)
	l.Logger.LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}
