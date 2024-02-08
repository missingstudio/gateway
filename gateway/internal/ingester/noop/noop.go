package noop

type NoOpIngester struct{}

func (n *NoOpIngester) Get(key string) ([]map[string]interface{}, error) {
	return nil, nil
}

func (n *NoOpIngester) Ingest(data map[string]interface{}, key string) {}

func (n *NoOpIngester) Close() error {
	return nil
}
