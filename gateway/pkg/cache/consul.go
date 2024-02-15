package cache

import (
	"context"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	api.Config
}

type ConsulClient struct {
	client *api.Client
}

func NewConsulClient(config *ConsulConfig) (ICache, error) {
	client, err := api.NewClient(&config.Config)
	if err != nil {
		return nil, err
	}

	return &ConsulClient{
		client: client,
	}, nil
}

func (c *ConsulClient) Get(ctx context.Context, key string) ([]byte, error) {
	kv, _, err := c.client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	var val []byte
	if kv != nil {
		val = kv.Value
	}

	return val, nil
}

func (c *ConsulClient) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	_, err := c.client.KV().Put(&api.KVPair{
		Key:   key,
		Value: value,
	}, nil)

	return err
}

func (c *ConsulClient) Delete(ctx context.Context, key string) error {
	_, err := c.client.KV().Delete(key, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConsulClient) IsAlive(context.Context) (bool, error) {
	health, _, err := c.client.Health().Checks("any", nil)
	return health.AggregatedStatus() == api.HealthPassing, err
}
