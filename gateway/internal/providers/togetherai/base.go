package togetherai

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.ProviderInterface = &togetherAIProvider{}

type togetherAIProvider struct {
	Name   string
	Config base.ProviderConfig
	TogetherAIHeaders
}

func NewTogetherAIProvider(headers TogetherAIHeaders) *togetherAIProvider {
	config := getTogetherAIConfig("https://api.together.xyz")

	return &togetherAIProvider{
		Name:              "TogetherAI",
		Config:            config,
		TogetherAIHeaders: headers,
	}
}

func (togetherAI togetherAIProvider) GetName() string {
	return togetherAI.Name
}

func (togetherAI togetherAIProvider) Validate() error {
	return utils.ValidateHeaders(togetherAI.TogetherAIHeaders)
}

func getTogetherAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
