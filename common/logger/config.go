package logger

type Config struct {
	Json bool `yaml:"json" json:"json,omitempty" mapstructure:"json" default:"false"`
}
