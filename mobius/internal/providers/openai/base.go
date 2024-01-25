package openai

import "github.com/missingstudio/studio/backend/internal/providers/base"

type OpenAIProvider struct {
	APIKey string
	Config base.ProviderConfig
}

func NewOpenAIProvider(token string, baseURL string) *OpenAIProvider {
	config := getOpenAIConfig(baseURL)
	return &OpenAIProvider{
		APIKey: token,
		Config: config,
	}
}

type OpenAIProviderFactory struct{}

func (f OpenAIProviderFactory) Create(token string) base.ProviderInterface {
	openAIProvider := NewOpenAIProvider(token, "https://api.openai.com")
	return openAIProvider
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
