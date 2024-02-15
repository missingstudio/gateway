package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConsulCache(t *testing.T) {
	config := Config{
		Driver:       Consul,
		ConsulConfig: ConsulConfig{},
	}
	reg, err := NewCache(&config)
	assert.NotNil(t, reg)
	assert.Nil(t, err)
}

func TestCreateRedisCache(t *testing.T) {
	config := Config{
		Driver:      Redis,
		RedisConfig: RedisConfig{},
	}
	reg, err := NewCache(&config)
	assert.Nil(t, reg)
	assert.NotNil(t, err)
}

func TestInvalidCacheDriver(t *testing.T) {
	config := Config{
		Driver: "invalid",
	}
	reg, err := NewCache(&config)
	assert.NotNil(t, err)
	assert.Nil(t, reg)
}
