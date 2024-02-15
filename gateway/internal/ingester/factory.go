package ingester

import (
	"context"
	"fmt"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/missingstudio/studio/backend/internal/ingester/influx"
	"github.com/sagikazarmark/slog-shim"
)

const (
	Influx = "influx"
	Nop    = "nop"
)

// NewIngester initializes the ingester instance based on Config
func NewIngester(ctx context.Context, cfg Config, logger *slog.Logger) (Ingester, error) {
	switch cfg.Provider {
	case Influx:
		// Create a new client using an InfluxDB server base URL and an authentication token
		client, err := influxdb3.New(influxdb3.ClientConfig{
			Host:         cfg.Influx.Host,
			Token:        cfg.Influx.Token,
			Organization: cfg.Influx.Organization,
			Database:     cfg.Influx.Database,
		})
		if err != nil {
			logger.Error("error starting influx server", "error", err)
			return nil, err
		}

		return influx.NewInfluxIngester(
			influx.WithClient(client),
			influx.WithLogger(logger),
			influx.WithDatabase(cfg.Influx.Database),
			influx.WithOrganization(cfg.Influx.Organization),
		), err

	default:
		return nil, fmt.Errorf("Unknown ingester Driver: %s", cfg.Provider)
	}
}

func GetIngesterWithDefault(ctx context.Context, cfg Config, logger *slog.Logger) Ingester {
	ingester, err := NewIngester(ctx, cfg, logger)
	if err != nil {
		ingester = NewNopIngester()
	}

	return ingester
}
