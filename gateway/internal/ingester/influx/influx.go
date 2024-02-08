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
	bucket       string
	Organization string
	logger       *slog.Logger
}

func NewInfluxDBIngester(host, token, organization, bucket string, logger *slog.Logger) (*InfluxDBIngester, error) {
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:         host,
		Token:        token,
		Organization: organization,
		Database:     bucket,
	})
	if err != nil {
		return nil, err
	}

	return &InfluxDBIngester{
		client:       client,
		Organization: organization,
		bucket:       bucket,
		logger:       logger,
	}, nil
}

func (in *InfluxDBIngester) Ingest(data map[string]interface{}, measurement string) {
	point := influxdb3.NewPoint(measurement, nil, data, time.Now())
	err := in.client.WritePoints(context.Background(), point)
	if err != nil {
		in.logger.Error("Not able to ingest into db", err)
	}
}

func (in *InfluxDBIngester) Get(query string) ([]map[string]interface{}, error) {
	result, err := in.client.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	data := []map[string]interface{}{}
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
