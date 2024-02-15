package ingester

type Ingester interface {
	Get(string) ([]map[string]interface{}, error)
	Ingest(map[string]interface{}, string)
	Close() error
}
