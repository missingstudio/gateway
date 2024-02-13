package cache

import (
	"time"

	"github.com/missingstudio/studio/backend/internal/encoding"
)

type store interface {
	Get(key string) ([]byte, error)
	Set(key string, value any, ttl time.Duration) error
}

type Cache struct {
	store store
}

func NewCache(s store) *Cache {
	return &Cache{
		store: s,
	}
}

func (c *Cache) ComputeHashKey(value string) string {
	return encoding.Encode(value)
}

func (c *Cache) SetValue(key string, value []byte, ttl time.Duration) error {
	return c.store.Set(c.ComputeHashKey(key), value, ttl)
}

func (c *Cache) GetValue(key string) ([]byte, error) {
	return c.store.Get(c.ComputeHashKey(key))
}
