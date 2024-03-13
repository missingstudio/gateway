package base

import (
	"context"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/provider"
)

type Config struct {
	BaseURL         string
	ChatCompletions string
}

type Info struct {
	Title       string
	Name        string
	Description string
}

type Provider interface {
	Info() Info
	Config() Config
	Models() []string
	Schema() []byte
}

var ProviderRegistry = map[string]func(provider.Provider) Provider{}

type ChatCompletionProvider interface {
	ChatCompletion(context.Context, *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error)
}
