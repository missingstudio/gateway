package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/missingstudio/ai/gateway/internal/api"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	"github.com/missingstudio/ai/gateway/internal/ratelimiter"

	"github.com/missingstudio/ai/gateway/pkg/database"

	"github.com/missingstudio/common/config"
	"github.com/missingstudio/common/logger"
)

type Config struct {
	Version       int                `yaml:"version"`
	App           api.Config         `yaml:"app,omitempty"`
	Log           logger.Config      `yaml:"log,omitempty"`
	Ingester      ingester.Config    `yaml:"ingester,omitempty"`
	Redis         RedisConfig        `yaml:"redis,omitempty"`
	Postgres      database.Config    `yaml:"postgres,omitempty"`
	Ratelimiter   ratelimiter.Config `yaml:"ratelimiter,omitempty"`
	EncryptionKey string             `yaml:"encryption_key" json:"encryption_key,omitempty" mapstructure:"encryption_key" default:""`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host,omitempty" mapstructure:"host" default:"localhost"`
	Port     int    `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"6379"`
	Username string `yaml:"username" mapstructure:"username" json:"username,omitempty"`
	Password string `yaml:"password" mapstructure:"password" json:"password,omitempty"`
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
