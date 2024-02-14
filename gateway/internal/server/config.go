package server

type Config struct {
	Host string `yaml:"host" json:"host,omitempty" mapstructure:"host" default:"0.0.0.0"`
	Port int    `yaml:"port" json:"port,omitempty" mapstructure:"port" default:"8080"`
}
