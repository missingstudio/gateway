package ratelimiter

import (
	"time"
)

type Rate struct {
	Period time.Duration
	Limit  int
}

func NewRate(duration, limit int) *Rate {
	return &Rate{
		Period: time.Duration(duration),
		Limit:  limit,
	}
}
