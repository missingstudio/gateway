package ingester

import (
	"context"

	"github.com/missingstudio/studio/backend/internal/ingester/influx"
	"github.com/missingstudio/studio/backend/internal/ingester/noop"
	"github.com/sagikazarmark/slog-shim"
)

type Ingester interface {
	Get(string) ([]map[string]interface{}, error)
	Ingest(map[string]interface{}, string)
	Close() error
}

func GetIngester(ctx context.Context, cfg Config, logger *slog.Logger) Ingester {
	if !cfg.Enabled {
		return &noop.NoOpIngester{}
	}

	switch cfg.Provider {
	case "influx":
		ingester, err := influx.NewInfluxDBIngester(cfg.Influx.Host, cfg.Influx.Token, cfg.Influx.Organization, "logs", logger)
		if err != nil {
			logger.Error("error starting influx server", "error", err)
			return nil
		}
		return ingester
	default:
		return &noop.NoOpIngester{}
	}
}
