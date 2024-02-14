package logger

import (
	"log/slog"
	"os"
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
