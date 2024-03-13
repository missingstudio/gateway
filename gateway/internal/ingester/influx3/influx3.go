package influx3

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

type InfluxDBIngester struct {
	client       *influxdb3.Client
	database     string
	organization string
	measurement  string
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
		measurement:  options.measurement,
		logger:       options.logger,
	}
}

func (in *InfluxDBIngester) Ingest(data map[string]any) {
	point := influxdb3.NewPoint(in.measurement, nil, data, time.Now())
	err := in.client.WritePoints(context.Background(), []*influxdb3.Point{point})
	if err != nil {
		in.logger.Error("Not able to ingest into db", err)
	}
}

func (in *InfluxDBIngester) Get() ([]map[string]any, error) {
	result, err := in.client.Query(context.Background(), fmt.Sprintf("select * from %s ORDER BY time DESC", in.measurement))
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
