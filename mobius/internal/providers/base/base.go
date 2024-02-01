package base

import (
	"context"
	"net/http"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}

type ProviderInterface interface {
	GetName() string
	Validate() error
}

type ChatCompletionInterface interface {
	ProviderInterface
	ChatCompletion(context.Context, []byte) (*http.Response, error)
}
