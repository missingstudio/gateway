package influx

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

type InfluxDBIngester struct {
	client       *influxdb3.Client
	database     string
	organization string
	logger       *slog.Logger
}

// NewOptions creates a new Options instance with provided functional options
func NewInfluxIngester(opts ...Option) *InfluxDBIngester {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return &InfluxDBIngester{
		client:       options.client,
		database:     options.database,
		organization: options.organization,
		logger:       options.logger,
	}
}

func (in *InfluxDBIngester) Ingest(data map[string]any, measurement string) {
	point := influxdb3.NewPoint(measurement, nil, data, time.Now())
	err := in.client.WritePoints(context.Background(), point)
	if err != nil {
		in.logger.Error("Not able to ingest into db", err)
	}
}

func (in *InfluxDBIngester) Get(query string) ([]map[string]any, error) {
	result, err := in.client.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	data := []map[string]any{}
	for result.Next() {
		data = append(data, result.Value())
	}

	return data, nil
}

func (in *InfluxDBIngester) Close() error {
	err := in.client.Close()
	if err != nil {
		log.Fatal("Error closing InfluxDB Client: " + err.Error())
		return err
	}

	return nil
}
