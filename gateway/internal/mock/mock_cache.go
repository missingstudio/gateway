package mock

import "time"

type MockStore struct {
	data map[string][]byte
}

func NewMockStore() *MockStore {
	return &MockStore{
		data: make(map[string][]byte),
	}
}

func (m *MockStore) Get(key string) ([]byte, error) {
	value, exists := m.data[key]
	if !exists {
		return nil, nil
	}
	return value, nil
}

func (m *MockStore) Set(key string, value any, ttl time.Duration) error {
	m.data[key] = value.([]byte)
	return nil
}
