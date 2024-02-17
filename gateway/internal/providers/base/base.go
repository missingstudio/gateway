package base

import (
	"context"
	"net/http"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}

type IProvider interface {
	Name() string
	Models() []string
	Schema() []byte
}

type ChatCompletionInterface interface {
	IProvider
	ChatCompletion(context.Context, []byte) (*http.Response, error)
}
