package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mcuadros/go-defaults"
	"github.com/missingstudio/studio/common/file"
	"gopkg.in/yaml.v2"
)

func Init(configFile string) error {
	if file.Exist(configFile) {
		return errors.New("config file already exists")
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

	if err := ioutil.WriteFile(configFile, data, 0o655); err != nil {
		return err
	}

	return nil
}
