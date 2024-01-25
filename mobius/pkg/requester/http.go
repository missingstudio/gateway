package requester

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultHTTPTimeout = 5

// HTTPClient represents a client to send HTTP requests.
type HTTPClient struct {
	client *http.Client

	// encoder is used to encode request bodies
	encoder Encoder

	// decoder is used to decode response bodies
	decoder Decoder

	// before is a function called before each
	// request is made. useful for like, auth sigs, etc.
	before func(*http.Request) error
}

// NewHTTPClient is used to build a new HTTPClient.
func NewHTTPClient(opts ...HTTPOption) *HTTPClient {
	return &HTTPClient{
		encoder: defaultEncoder,
		decoder: defaultDecoder,
		client: &http.Client{
			Timeout: time.Second * time.Duration(defaultHTTPTimeout),
		},
		before: func(_ *http.Request) error { return nil },
	}
}

func (c *HTTPClient) Do(req *http.Request) ([]byte, error) {
	if err := c.before(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if int(resp.StatusCode/100) != 2 {
		return nil, fmt.Errorf("http status not 2xx: %d %s", resp.StatusCode, string(body))
	}
	return body, nil
}
