package cache

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Database int32
	Password string
}

type RedisClient struct {
	client *redis.Client
	rt     time.Duration
	wt     time.Duration
}

type RedisOption func(*RedisClient)

// WithReadTimeout set read timeout
func WithReadTimeout(t time.Duration) RedisOption {
	return func(o *RedisClient) {
		o.wt = t
	}
}

// WithWriteTimeout set write timeout
func WithWriteTimeout(t time.Duration) RedisOption {
	return func(o *RedisClient) {
		o.rt = t
	}
}

func NewRedisClient(config *RedisConfig, opts ...RedisOption) (ICache, error) {
	options := &redis.Options{
		Addr:     net.JoinHostPort(config.Host, config.Port),
		DB:       int(config.Database),
		Password: config.Password,
	}

	client := redis.NewClient(options)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	rc := &RedisClient{
		client: client,
	}

	for _, o := range opts {
		o(rc)
	}

	return rc, nil
}

func (rc *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (rc *RedisClient) MGet(ctx context.Context, keys ...string) ([]any, error) {
	return rc.client.MGet(ctx, keys...).Result()
}

func (rc *RedisClient) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	err := rc.client.Set(ctx, key, string(value), ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) Delete(ctx context.Context, key string) error {
	err := rc.client.Del(ctx, key).Err()
	return err
}

func (rc *RedisClient) IsAlive(ctx context.Context) (bool, error) {
	ping, err := rc.client.Ping(ctx).Result()
	if err != nil {
		return false, err
	}

	if ping == "PONG" {
		return true, err
	}

	return false, err
}

func (rc *RedisClient) Disconnect() error {
	err := rc.client.Close()
	if err != nil {
		return err
	}

	return nil
}
