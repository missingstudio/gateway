package ingester

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/missingstudio/studio/backend/internal/ingester/clickhouse"
	"github.com/missingstudio/studio/backend/internal/ingester/influx3"
	"github.com/sagikazarmark/slog-shim"
)

const (
	Influx3    = "influx3"
	ClickHouse = "clickhouse"
	Nop        = "nop"
)

// NewIngester initializes the ingester instance based on Config
func NewIngester(ctx context.Context, cfg Config, logger *slog.Logger) (Ingester, error) {
	switch cfg.Provider {
	case Influx3:
		// Create a new client using an InfluxDB server base URL and an authentication token
		client, err := influxdb3.New(influxdb3.ClientConfig{
			Host:         cfg.Influx3.Host,
			Token:        cfg.Influx3.Token,
			Organization: cfg.Influx3.Organization,
			Database:     cfg.Influx3.Database,
		})
		if err != nil {
			logger.Error("error starting influx server", "error", err)
			return nil, err
		}

		return influx3.NewInfluxIngester(
			influx3.WithClient(client),
			influx3.WithLogger(logger),
			influx3.WithDatabase(cfg.Influx3.Database),
			influx3.WithOrganization(cfg.Influx3.Organization),
			influx3.WithMeasurement(cfg.Influx3.Measurement),
		), err
	case ClickHouse:
		db, err := sql.Open("clickhouse", cfg.Clickhouse.URL)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return clickhouse.NewClickHouseIngester(db, cfg.Clickhouse.Table)
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
