package openai

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/core/connection"
	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionInterface = &openAIProvider{}

type openAIProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   connection.Connection
}

func (anyscale openAIProvider) Info() base.ProviderInfo {
	return anyscale.info
}

func (oai openAIProvider) Config() base.ProviderConfig {
	return oai.config
}

func (oai openAIProvider) Schema() []byte {
	return schema
}

func getOpenAIInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title:       "OpenAI",
		Name:        "openai",
		Description: `OpenAI API platform offers latest models and guides for safety best practices.`,
	}
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["openai"] = func(connection connection.Connection) base.IProvider {
		config := getOpenAIConfig("https://api.openai.com")
		return &openAIProvider{
			info:   getOpenAIInfo(),
			config: config,
			conn:   connection,
		}
	}
}
