package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/missingstudio/studio/common/config"
)

type Config struct {
	Port          int  `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"8080"`
	LogFormatJson bool `yaml:"log_format_json" json:"log_format_json,omitempty" mapstructure:"log_format_json" default:"false"`
}

func Load(serverConfigFileFromFlag string) (*Config, error) {
	conf := &Config{}

	var options []config.LoaderOption
	options = append(options, config.WithName("config"))
	options = append(options, config.WithEnvKeyReplacer(".", "_"))
	options = append(options, config.WithEnvPrefix("MOBIUS"))
	if p, err := os.Getwd(); err == nil {
		options = append(options, config.WithPath(p))
	}
	if execPath, err := os.Executable(); err == nil {
		options = append(options, config.WithPath(filepath.Dir(execPath)))
	}
	if currentHomeDir, err := os.UserHomeDir(); err == nil {
		options = append(options, config.WithPath(currentHomeDir))
		options = append(options, config.WithPath(filepath.Join(currentHomeDir, ".config")))
	}

	// override all config sources and prioritize one from file
	if serverConfigFileFromFlag != "" {
		options = append(options, config.WithFile(serverConfigFileFromFlag))
	}

	l := config.NewLoader(options...)
	if err := l.Load(conf); err != nil {
		if !errors.As(err, &config.ConfigFileNotFoundError{}) {
			return nil, err
		}
	}

	return conf, nil
}
