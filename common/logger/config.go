package logger

type Config struct {
	Level int  `yaml:"level" json:"level,omitempty" mapstructure:"level" default:"0"`
	Json  bool `yaml:"json" json:"json,omitempty" mapstructure:"json" default:"false"`
}
