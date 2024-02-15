package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
	"github.com/missingstudio/studio/backend/internal/server"

	"github.com/missingstudio/studio/common/config"
	"github.com/missingstudio/studio/common/logger"
)

type Config struct {
	App         server.Config      `yaml:"app,omitempty"`
	Log         logger.Config      `yaml:"log,omitempty"`
	Ingester    ingester.Config    `yaml:"ingester,omitempty"`
	Redis       RedisConfig        `yaml:"redis,omitempty"`
	Ratelimiter ratelimiter.Config `yaml:"ratelimiter,omitempty"`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host,omitempty" mapstructure:"host" default:"localhost"`
	Port     int    `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"6379"`
	Username string `yaml:"username" mapstructure:"username" json:"username,omitempty"`
	Password string `yaml:"password" mapstructure:"password" json:"password,omitempty"`
	Database string `yaml:"database" mapstructure:"database"  json:"database,omitempty"`
}

func Load(serverConfigFileFromFlag string) (*Config, error) {
	conf := &Config{}

	var options []config.LoaderOption
	options = append(options, config.WithName("config"))
	options = append(options, config.WithEnvKeyReplacer(".", "_"))
	options = append(options, config.WithEnvPrefix("gateway"))
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
