package cache_test

import (
	"testing"
	"time"

	"github.com/missingstudio/studio/backend/internal/cache"
	"github.com/missingstudio/studio/backend/internal/mock"
)

func TestCache_SetValue(t *testing.T) {
	// Create a new instance of the MockStore
	mockStore := mock.NewMockStore()

	// Create a new Cache with the MockStore
	cache := cache.NewCache(mockStore)

	key := "key"
	value := []byte("value")
	ttl := 5 * time.Second

	// Set the value in the cache
	err := cache.SetValue(key, value, ttl)
	if err != nil {
		t.Errorf("Error setting value in cache: %v", err)
	}

	// Verify that the value is correctly stored in the MockStore
	storedValue, err := mockStore.Get(cache.ComputeHashKey(key))
	if err != nil {
		t.Errorf("Error getting value from store: %v", err)
	}

	if storedValue == nil {
		t.Error("Expected value in store, but got nil")
	} else if string(storedValue) != string(value) {
		t.Errorf("Expected value '%s' in store, but got '%s'", value, storedValue)
	}
}
