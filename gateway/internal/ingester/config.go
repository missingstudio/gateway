package ingester

type Config struct {
	Enabled  bool          `yaml:"enabled" mapstructure:"enabled" json:"enabled,omitempty"`
	Provider string        `yaml:"provider" mapstructure:"provider" json:"provider,omitempty"`
	Influx3  Influx3Config `yaml:"influx3" mapstructure:"influx3" json:"influx3,omitempty"`
}

type Influx3Config struct {
	Host         string `yaml:"host" mapstructure:"host" default:"http://localhost:1234" json:"host,omitempty"`
	Token        string `yaml:"token" mapstructure:"token" json:"token,omitempty"`
	Organization string `yaml:"organization" mapstructure:"organization"  json:"organization,omitempty"`
	Database     string `yaml:"database" mapstructure:"database"  json:"database,omitempty"`
	Measurement  string `yaml:"measurement" mapstructure:"measurement"  json:"measurement,omitempty" default:"analytics"`
}
