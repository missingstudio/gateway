package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mcuadros/go-defaults"
	"github.com/missingstudio/common/file"
	"gopkg.in/yaml.v2"
)

var ErrConfigFileExit = errors.New("config file already exists")

func Init(configFile string) error {
	if file.Exist(configFile) {
		return ErrConfigFileExit
	}

	cfg := &Config{}
	defaults.SetDefaults(cfg)

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if !file.DirExists(configFile) {
			_ = os.MkdirAll(filepath.Dir(configFile), 0o755)
		}
	}

	if err := os.WriteFile(configFile, data, 0o655); err != nil {
		return err
	}

	return nil
}
