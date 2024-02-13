package redis

import (
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Options struct {
	client *redis.Client
	logger *slog.Logger
	rt     time.Duration
	wt     time.Duration
}

type Option func(*Options)

// WithClient configures Redis transport to use the given client
func WithClient(c *redis.Client) Option {
	return func(o *Options) {
		o.client = c
	}
}

// WithLogger set logger
func WithLogger(c *slog.Logger) Option {
	return func(o *Options) {
		o.logger = c
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.wt = t
	}
}

// WithWriteTimeout set write timeout
func WithWriteTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.rt = t
	}
}
