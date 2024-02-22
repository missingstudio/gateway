package ingester

var _ Ingester = &nopIngester{}

type nopIngester struct{}

func NewNopIngester() *nopIngester {
	return &nopIngester{}
}

func (n *nopIngester) Get(key string) ([]map[string]any, error) {
	return nil, nil
}
func (n *nopIngester) Ingest(data map[string]any, key string) {}

func (n *nopIngester) Close() error {
	return nil
}
