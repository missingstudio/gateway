package openai

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
)

type OpenAIProvider struct {
	APIKey string
	Config base.ProviderConfig
}

func NewOpenAIProvider(apikey string, baseURL string) *OpenAIProvider {
	config := getOpenAIConfig(baseURL)

	return &OpenAIProvider{
		APIKey: apikey,
		Config: config,
	}
}

type OpenAIProviderFactory struct{}

func (f OpenAIProviderFactory) Create(apikey string) base.ProviderInterface {
	openAIProvider := NewOpenAIProvider(apikey, "https://api.openai.com")
	return openAIProvider
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
