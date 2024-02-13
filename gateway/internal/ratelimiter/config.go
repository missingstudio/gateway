package ratelimiter

type Config struct {
	Type             string `yaml:"type" json:"type,omitempty" mapstructure:"type" default:"sliding_window"`
	DurationInSecond int    `yaml:"duration_in_seconds" json:"duration_in_seconds,omitempty" mapstructure:"duration_in_seconds" default:"1"`
	NumberOfRequests int    `yaml:"number_of_requests" json:"number_of_requests,omitempty" mapstructure:"number_of_requests" default:"20"`
}
