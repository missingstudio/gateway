package ingester

type Ingester interface {
	Get() ([]map[string]any, error)
	Ingest(map[string]any)
	Close() error
}
