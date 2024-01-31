package config

import (
	"errors"
	"os"
	"path/filepath"

	msconfig "github.com/missingstudio/studio/common/config"
)

type Config struct {
	Port          int  `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"8080"`
	LogFormatJson bool `yaml:"log_format_json" json:"log_format_json,omitempty" mapstructure:"log_format_json" default:"false"`
}

func Load(serverConfigFileFromFlag string) (*Config, error) {
	conf := &Config{}

	var options []msconfig.LoaderOption
	options = append(options, msconfig.WithName("config"))
	options = append(options, msconfig.WithEnvKeyReplacer(".", "_"))
	options = append(options, msconfig.WithEnvPrefix("MOBIUS"))
	if p, err := os.Getwd(); err == nil {
		options = append(options, msconfig.WithPath(p))
	}
	if execPath, err := os.Executable(); err == nil {
		options = append(options, msconfig.WithPath(filepath.Dir(execPath)))
	}
	if currentHomeDir, err := os.UserHomeDir(); err == nil {
		options = append(options, msconfig.WithPath(currentHomeDir))
		options = append(options, msconfig.WithPath(filepath.Join(currentHomeDir, ".config")))
	}

	// override all config sources and prioritize one from file
	if serverConfigFileFromFlag != "" {
		options = append(options, msconfig.WithFile(serverConfigFileFromFlag))
	}

	l := msconfig.NewLoader(options...)
	if err := l.Load(conf); err != nil {
		if !errors.As(err, &msconfig.ConfigFileNotFoundError{}) {
			return nil, err
		}
	}

	return conf, nil
}
