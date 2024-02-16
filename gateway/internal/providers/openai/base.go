package openai

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.ProviderInterface = &openAIProvider{}

type openAIProvider struct {
	Name   string
	Config base.ProviderConfig
	OpenAIHeaders
}

func NewOpenAIProvider(headers OpenAIHeaders) *openAIProvider {
	config := getOpenAIConfig("https://api.openai.com")

	return &openAIProvider{
		Name:          "OpenAI",
		Config:        config,
		OpenAIHeaders: headers,
	}
}

func (oai openAIProvider) GetName() string {
	return oai.Name
}

func (oai openAIProvider) Validate() error {
	return utils.ValidateHeaders(oai.OpenAIHeaders)
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
