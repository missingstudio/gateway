package requester

import (
	"bytes"
	"encoding/json"
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

func (c *HTTPClient) SendRequest(req *http.Request, response any, outputResponse bool) (*http.Response, error) {
	if err := c.before(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if !outputResponse {
		defer resp.Body.Close()
	}

	if c.IsFailureStatusCode(resp) {
		return nil, fmt.Errorf("http status not 2xx: %d", resp.StatusCode)
	}

	if outputResponse {
		var buf bytes.Buffer
		tee := io.TeeReader(resp.Body, &buf)
		err = DecodeResponse(tee, response)

		resp.Body = io.NopCloser(&buf)
	} else {
		err = json.NewDecoder(resp.Body).Decode(response)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HTTPClient) SendRequestRaw(req *http.Request) (*http.Response, error) {
	if err := c.before(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.IsFailureStatusCode(resp) {
		return nil, fmt.Errorf("http status not 2xx: %d", resp.StatusCode)
	}

	return resp, nil
}

func (r *HTTPClient) IsFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

type Stringer interface {
	GetString() *string
}

func DecodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return DecodeString(body, result)
	}

	if stringer, ok := v.(Stringer); ok {
		return DecodeString(body, stringer.GetString())
	}

	return json.NewDecoder(body).Decode(v)
}

func DecodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}
