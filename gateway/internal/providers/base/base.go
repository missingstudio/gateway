package base

import (
	"context"
	"net/http"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}
type ProviderInfo struct {
	Title       string
	Name        string
	Description string
}

type IProvider interface {
	Info() ProviderInfo
	Config() ProviderConfig
	Models() []string
	Schema() []byte
}

type ChatCompletionInterface interface {
	IProvider
	ChatCompletion(context.Context, []byte) (*http.Response, error)
}
