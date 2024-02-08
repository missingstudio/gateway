package ingester

import (
	"context"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/missingstudio/studio/backend/internal/ingester/influx"
	"github.com/sagikazarmark/slog-shim"
)

type Ingester interface {
	Get(string) ([]map[string]interface{}, error)
	Ingest(map[string]interface{}, string)
	Close() error
}

func GetIngester(ctx context.Context, cfg Config, logger *slog.Logger) Ingester {
	if !cfg.Enabled {
		return nil
	}

	switch cfg.Provider {
	case "influx":
		// Create a new client using an InfluxDB server base URL and an authentication token
		client, err := influxdb3.New(influxdb3.ClientConfig{
			Host:         cfg.Influx.Host,
			Token:        cfg.Influx.Token,
			Organization: cfg.Influx.Organization,
			Database:     cfg.Influx.Database,
		})
		if err != nil {
			logger.Error("error starting influx server", "error", err)
			return nil
		}

		return influx.NewInfluxIngester(
			influx.WithClient(client),
			influx.WithLogger(logger),
			influx.WithDatabase(cfg.Influx.Database),
			influx.WithOrganization(cfg.Influx.Organization),
		)
	default:
		return nil
	}
}

func GetIngesterWithDefault(ctx context.Context, cfg Config, logger *slog.Logger) Ingester {
	ingester := GetIngester(ctx, cfg, logger)
	if ingester == nil {
		ingester = &NopIngester{}
	}

	return ingester
}
