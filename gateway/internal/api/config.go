package api

import "time"

type Config struct {
	Host           string               `yaml:"host" json:"host,omitempty" mapstructure:"host" default:"0.0.0.0"`
	Port           int                  `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"8080"`
	RequestTimeout time.Duration        `yaml:"request_timeout" json:"request_timeout,omitempty" mapstructure:"request_timeout" default:"5s"`
	Authentication AuthenticationConfig `yaml:"authentication" mapstructure:"authentication"`
}

type AuthenticationConfig struct {
	Enabled bool `yaml:"enabled" json:"enabled,omitempty" mapstructure:"enabled" default:"false"`
}
