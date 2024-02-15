package cache

type Config struct {
	Driver       string
	RedisConfig  RedisConfig
	ConsulConfig ConsulConfig
}
