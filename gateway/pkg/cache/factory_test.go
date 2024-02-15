package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConsulCache(t *testing.T) {
	config := Config{
		Provider:     Consul,
		ConsulConfig: ConsulConfig{},
	}
	reg, err := NewCache(&config)
	assert.NotNil(t, reg)
	assert.Nil(t, err)
}

func TestCreateRedisCache(t *testing.T) {
	config := Config{
		Provider:    Redis,
		RedisConfig: RedisConfig{},
	}
	reg, err := NewCache(&config)
	assert.Nil(t, reg)
	assert.NotNil(t, err)
}

func TestInvalidCacheProvider(t *testing.T) {
	config := Config{
		Provider: "invalid",
	}
	reg, err := NewCache(&config)
	assert.NotNil(t, err)
	assert.Nil(t, reg)
}
