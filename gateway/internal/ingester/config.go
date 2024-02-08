package ingester

type Config struct {
	Enabled  bool         `yaml:"enabled" mapstructure:"enabled" json:"enabled,omitempty"`
	Provider string       `yaml:"provider" mapstructure:"provider" json:"provider,omitempty"`
	Influx   InfluxConfig `yaml:"influx" mapstructure:"influx" json:"influx,omitempty"`
}

type InfluxConfig struct {
	Host         string `yaml:"host" mapstructure:"host" default:"none" json:"host,omitempty"`
	Token        string `yaml:"token" mapstructure:"token" default:"json" json:"token,omitempty"`
	Organization string `yaml:"organization" mapstructure:"organization" default:"json" json:"organization,omitempty"`
	Database     string `yaml:"database" mapstructure:"database" default:"json" json:"database,omitempty"`
}
