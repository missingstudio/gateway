package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"strings"

	"github.com/jeremywohl/flatten"
	"github.com/mcuadros/go-defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ConfigFileNotFoundError is returned when the config file is not found
// Viper will load from env or defaults
type ConfigFileNotFoundError struct {
	err error
}

func (err ConfigFileNotFoundError) Error() string {
	return fmt.Sprintf("unable to find config file, loading from env and defaults: %v", err.err)
}

func (err *ConfigFileNotFoundError) Unwrap() error {
	return err.err
}

type Loader struct {
	v    *viper.Viper
	opts []viper.DecoderConfigOption
}

type LoaderOption func(*Loader)

// WithViper sets the given viper instance for loading configs
// instead of the default configured one
func WithViper(in *viper.Viper) LoaderOption {
	return func(l *Loader) {
		l.v = in
	}
}

// WithFile explicitly defines the path, name and extension
// of the config file
func WithFile(file string) LoaderOption {
	return func(l *Loader) {
		l.v.SetConfigFile(file)
	}
}

// WithName sets the file name of the config file without
// the extension
func WithName(in string) LoaderOption {
	return func(l *Loader) {
		l.v.SetConfigName(in)
	}
}

// WithPath adds config path to search the config file in,
// can be used multiple times to add multiple paths to search
func WithPath(in string) LoaderOption {
	return func(l *Loader) {
		l.v.AddConfigPath(in)
	}
}

// WithType sets the type of the configuration e.g. "json",
// "yaml", "hcl"
// Also used for the extension of the file
func WithType(in string) LoaderOption {
	return func(l *Loader) {
		l.v.SetConfigType(in)
	}
}

// WithBindPFlags binds viper to pflags based on the
// tags in the struct. Use tag `cmdx` to bind struct
// field to cli flag.
// e.g.
//
//	type Config struct {
//		Host string `yaml:"host" cmdx:"host"`
//	}
func WithBindPFlags(pfs *pflag.FlagSet, cfg any) LoaderOption {
	return func(l *Loader) {
		reflectedCfg := reflect.TypeOf(cfg).Elem()

		flagTags := []string{}
		for i := 0; i < reflectedCfg.NumField(); i++ {
			if tag := reflectedCfg.Field(i).Tag.Get("cmdx"); tag != "" {
				flagTags = append(flagTags, tag)
			}
		}

		if pfs != nil {
			for _, tag := range flagTags {
				l.v.BindPFlag(tag, pfs.Lookup(tag))
			}
		}
	}
}

// WithEnvPrefix sets the prefix for keys when checking for configs
// in environment variables. Internally concatenates with keys
// with `_` in between
func WithEnvPrefix(in string) LoaderOption {
	return func(l *Loader) {
		l.v.SetEnvPrefix(in)
	}
}

// WithEnvKeyReplacer sets the `old` string to be replaced with
// the `new` string environmental variable to a key that does
// not match it.
func WithEnvKeyReplacer(old string, new string) LoaderOption {
	return func(l *Loader) {
		l.v.SetEnvKeyReplacer(strings.NewReplacer(old, new))
	}
}

// WithDecoderConfigOption sets the decoder config options for
// viper. See https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig
// for more details
func WithDecoderConfigOption(opts ...viper.DecoderConfigOption) LoaderOption {
	return func(l *Loader) {
		l.opts = append(l.opts, opts...)
	}
}

// StringToJsonFunc is a mapstructure.DecodeHookFunc that converts a string to an any
// if the string is valid json. This is useful for unmarshalling json strings into a map.
// For example, if you have a struct with a field of type map[string]string like labels or annotations,
func StringToJsonFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data any) (any, error) {
		// if type is string and can be parsed as json, parse it
		if f.Kind() == reflect.String && t.Kind() == reflect.Map {
			if f.Kind() == reflect.String && t.Kind() == reflect.Map {
				var m map[string]any
				if err := json.Unmarshal([]byte(data.(string)), &m); err == nil {
					return m, nil
				}
			}
		}
		return data, nil
	}
}

// NewLoader returns a config loader with given LoaderOption(s)
func NewLoader(options ...LoaderOption) *Loader {
	loader := &Loader{
		v: getViperWithDefaults(),
		opts: []viper.DecoderConfigOption{
			viper.DecodeHook(
				mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToSliceHookFunc(","), // default delimiter
					StringToJsonFunc(),
				),
			),
		},
	}

	for _, option := range options {
		option(loader)
	}
	return loader
}

// Load loads configuration into the given mapstructure (https://github.com/mitchellh/mapstructure)
// from a config.yaml file and overrides with any values set in env variables
func (l *Loader) Load(config any) error {
	if err := verifyParamIsPtrToStructElsePanic(config); err != nil {
		return err
	}

	l.v.AutomaticEnv()

	var werr error

	if err := l.v.ReadInConfig(); err != nil {
		pathErr := new(fs.PathError)
		if errors.As(err, &pathErr) || errors.As(err, &viper.ConfigFileNotFoundError{}) {
			werr = ConfigFileNotFoundError{err}
		} else {
			return fmt.Errorf("unable to read config file: %w", err)
		}
	}

	configKeys, err := getFlattenedStructKeys(config)
	if err != nil {
		return fmt.Errorf("unable to get all config keys from struct: %v", err)
	}

	// Bind each conf fields from struct to environment vars
	for key := range configKeys {
		if err := l.v.BindEnv(configKeys[key]); err != nil {
			return fmt.Errorf("unable to bind env keys: %v", err)
		}
	}

	// set defaults using the default struct tag
	defaults.SetDefaults(config)

	if err := l.v.Unmarshal(config, l.opts...); err != nil {
		return fmt.Errorf("unable to load config to struct: %v", err)
	}

	if werr != nil {
		return werr
	}

	return nil
}

func verifyParamIsPtrToStructElsePanic(param any) error {
	value := reflect.ValueOf(param)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("require ptr to a struct for Load. Got %v", value.Kind())
	} else {
		value = reflect.Indirect(value)
		if value.Kind() != reflect.Struct {
			return fmt.Errorf("require ptr to a struct for Load. got ptr to %v", value.Kind())
		}
	}
	return nil
}

func getViperWithDefaults() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}

func getFlattenedStructKeys(config any) ([]string, error) {
	var structMap map[string]any
	if err := mapstructure.Decode(config, &structMap); err != nil {
		return nil, err
	}

	flat, err := flatten.Flatten(structMap, "", flatten.DotStyle)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(flat))
	for k := range flat {
		keys = append(keys, k)
	}

	return keys, nil
}
