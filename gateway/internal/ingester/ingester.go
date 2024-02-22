package ingester

type Ingester interface {
	Get(string) ([]map[string]any, error)
	Ingest(map[string]any, string)
	Close() error
}
