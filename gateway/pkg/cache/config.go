package cache

type Config struct {
	Provider     string
	RedisConfig  RedisConfig
	ConsulConfig ConsulConfig
}
