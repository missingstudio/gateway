package requester

import (
	"net/http"
)

type HTTPOption func(*HTTPClient)

func WithClient(c *http.Client) HTTPOption {
	return func(client *HTTPClient) {
		client.client = c
	}
}

func WithEncoder(fn func(obj any) ([]byte, error)) HTTPOption {
	return func(c *HTTPClient) {
		c.encoder = fn
	}
}

func WithDecoder(fn func([]byte, any) error) HTTPOption {
	return func(c *HTTPClient) {
		c.decoder = fn
	}
}

func WithBefore(fn func(*http.Request) error) HTTPOption {
	return func(c *HTTPClient) {
		c.before = fn
	}
}
