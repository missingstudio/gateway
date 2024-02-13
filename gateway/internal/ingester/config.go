package ingester

type Config struct {
	Enabled  bool         `yaml:"enabled" mapstructure:"enabled" json:"enabled,omitempty"`
	Provider string       `yaml:"provider" mapstructure:"provider" json:"provider,omitempty"`
	Influx   InfluxConfig `yaml:"influx" mapstructure:"influx" json:"influx,omitempty"`
}

type InfluxConfig struct {
	Host         string `yaml:"host" mapstructure:"host" default:"http://localhost:1234" json:"host,omitempty"`
	Token        string `yaml:"token" mapstructure:"token" json:"token,omitempty"`
	Organization string `yaml:"organization" mapstructure:"organization"  json:"organization,omitempty"`
	Database     string `yaml:"database" mapstructure:"database"  json:"database,omitempty"`
}
