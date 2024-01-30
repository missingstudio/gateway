package config

type Config struct {
	Port          int  `mapstructure:"port" default:"8080"`
	LogFormatJson bool `mapstructure:"log_format_json" default:"false"`
}
