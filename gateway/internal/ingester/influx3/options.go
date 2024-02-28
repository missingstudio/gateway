package influx3

import (
	"log/slog"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

type Options struct {
	client       *influxdb3.Client
	database     string
	organization string
	measurement  string
	logger       *slog.Logger
}

type Option func(*Options)

// WithClient sets the InfluxDB client in Options
func WithClient(client *influxdb3.Client) Option {
	return func(o *Options) {
		o.client = client
	}
}

// WithDatabase sets the database in Options
func WithDatabase(database string) Option {
	return func(o *Options) {
		o.database = database
	}
}

// WithOrganization sets the organization in Options
func WithOrganization(organization string) Option {
	return func(o *Options) {
		o.organization = organization
	}
}

// WithMeasurement sets the measurement in Options
func WithMeasurement(measurement string) Option {
	return func(o *Options) {
		o.measurement = measurement
	}
}

// WithLogger sets the logger in Options
func WithLogger(logger *slog.Logger) Option {
	return func(o *Options) {
		o.logger = logger
	}
}
