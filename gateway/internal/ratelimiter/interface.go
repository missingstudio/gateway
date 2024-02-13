package ratelimiter

type RateLimiterProvider interface {
	Validate(key string) bool
}
