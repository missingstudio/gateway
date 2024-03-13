package openai

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &openAIProvider{}

type openAIProvider struct {
	info     base.Info
	config   base.Config
	provider provider.Provider
}

func (anyscale openAIProvider) Info() base.Info {
	return anyscale.info
}

func (oai openAIProvider) Config() base.Config {
	return oai.config
}

func (oai openAIProvider) Schema() []byte {
	return schema
}

func getOpenAIInfo() base.Info {
	return base.Info{
		Title:       "OpenAI",
		Name:        "openai",
		Description: `OpenAI API platform offers latest models and guides for safety best practices.`,
	}
}

func getOpenAIConfig(baseURL string) base.Config {
	return base.Config{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["openai"] = func(provider provider.Provider) base.Provider {
		config := getOpenAIConfig("https://api.openai.com")
		return &openAIProvider{
			info:     getOpenAIInfo(),
			config:   config,
			provider: provider,
		}
	}
}
