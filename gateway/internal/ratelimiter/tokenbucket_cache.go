package ratelimiter

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type TokenBucketCache interface {
	Refill(keyname string, value int)
	Validate(keyname string, bucket_size int) (count int, isCreated bool)
}

type redisTokenBucketCache struct {
	logger *slog.Logger
	client *redis.Client
}

func NewRedisTokenBucketCache(logger *slog.Logger, client *redis.Client) TokenBucketCache {
	return &redisTokenBucketCache{
		logger: logger,
		client: client,
	}
}

func (c *redisTokenBucketCache) Refill(keyname string, bucketSize int) {
	script := redis.NewScript(`
		local keyname = KEYS[1]
		local bucket_size = ARGV[1]
		redis.call('SET', keyname, bucket_size)
		return true
	`)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := script.Run(ctx, c.client, []string{keyname}, bucketSize).Result()
	if err != nil {
		c.logger.Error("failed to refill bucket")
	}
}

func (c *redisTokenBucketCache) Validate(keyname string, bucket_size int) (count int, isCreated bool) {
	script := redis.NewScript(`
		local keyname = KEYS[1]
		local bucket_size = ARGV[1]
		local count = redis.call('GET', keyname)
		local is_created = false
		if (count==false) then
			redis.call('SET', keyname, bucket_size)
			count = bucket_size
			is_created = true
		end
		if tonumber(count) > 0 then
			redis.call('DECR', keyname)
		end
		return {count, is_created}
	`)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n, err := script.Run(ctx, c.client, []string{keyname}, bucket_size).Result()
	if err != nil {
		c.logger.Error("failed to validate rate limiter")
		return
	}

	var ok bool
	var arr []any
	if arr, ok = n.([]any); !ok {
		c.logger.Error("failed to parse redis result")
		return
	}

	if val, ok := arr[0].(string); ok {
		ival, err := strconv.Atoi(val)
		if err != nil {
			c.logger.Error("failed to parse bucket count")
			return
		}
		count = ival
	}

	if val, ok := arr[1].(int64); ok {
		isCreated = (val == 1)
	}
	return
}
