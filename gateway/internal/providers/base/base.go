package base

import (
	"context"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/connection"
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

// ProviderRegistry holds all supported provider for which connections
// can be initialized
var ProviderRegistry = map[string]func(connection.Connection) IProvider{}

type ChatCompletionInterface interface {
	IProvider
	ChatCompletion(context.Context, *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error)
}
