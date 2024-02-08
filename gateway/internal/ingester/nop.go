package ingester

type NopIngester struct{}

func (n *NopIngester) Get(key string) ([]map[string]interface{}, error) {
	return nil, nil
}
func (n *NopIngester) Ingest(data map[string]interface{}, key string) {}
func (n *NopIngester) Close() error {
	return nil
}
