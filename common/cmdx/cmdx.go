package cmdx

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mcuadros/go-defaults"
	"github.com/missingstudio/common/config"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const (
	STUDIO_CONFIG_DIR = "STUDIO_CONFIG_DIR"
	XDG_CONFIG_HOME   = "XDG_CONFIG_HOME"
	APP_DATA          = "AppData"
	LOCAL_APP_DATA    = "LocalAppData"
)

var ErrConfigFileExit = errors.New("config file already exists")

type Config struct {
	filename      string
	boundedPFlags *pflag.FlagSet
}

type ConfigLoaderOpt func(c *Config)

func WithFlags(pfs *pflag.FlagSet) ConfigLoaderOpt {
	return func(c *Config) {
		c.boundedPFlags = pfs
	}
}

// SetConfig allows to set a client config file. It is used to
// load and save a config file for command line clients.
func SetConfig(app string) *Config {
	return &Config{
		filename: configFile(app),
	}
}

func (c *Config) File() string {
	return c.filename
}

func (c *Config) Defaults(cfg any) {
	defaults.SetDefaults(cfg)
}

func (c *Config) Init(cfg any) error {
	defaults.SetDefaults(cfg)

	if fileExist(c.filename) {
		return ErrConfigFileExit
	}

	return c.Write(cfg)
}

func (c *Config) Read() (string, error) {
	cfg, err := os.ReadFile(c.filename)
	return string(cfg), err
}

func (c *Config) Write(cfg any) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(c.filename); os.IsNotExist(err) {
		_ = os.MkdirAll(configDir("mstudio"), 0o700)
	}

	if err := os.WriteFile(c.filename, data, 0o655); err != nil {
		return err
	}
	return nil
}

func (c *Config) Load(cfg any, opts ...ConfigLoaderOpt) error {
	for _, opt := range opts {
		opt(c)
	}

	loaderOpts := []config.LoaderOption{config.WithFile(c.filename)}

	if c.boundedPFlags != nil {
		loaderOpts = append(loaderOpts, config.WithBindPFlags(c.boundedPFlags, cfg))
	}

	loader := config.NewLoader(loaderOpts...)

	if err := loader.Load(cfg); err != nil {
		return err
	}
	return nil
}

func configFile(app string) string {
	file := app + ".yml"
	return filepath.Join(configDir("mstudio"), file)
}

func configDir(root string) string {
	var path string
	if a := os.Getenv(STUDIO_CONFIG_DIR); a != "" {
		path = a
	} else if b := os.Getenv(XDG_CONFIG_HOME); b != "" {
		path = filepath.Join(b, root)
	} else if c := os.Getenv(APP_DATA); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, root)
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, ".config", root)
	}

	if !dirExists(path) {
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
	}

	return path
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func dirExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}
