package togetherai

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.IProvider = &togetherAIProvider{}

type togetherAIProvider struct {
	name   string
	config base.ProviderConfig
	TogetherAIHeaders
}

func NewTogetherAIProvider(headers TogetherAIHeaders) *togetherAIProvider {
	config := getTogetherAIConfig("https://api.together.xyz")

	return &togetherAIProvider{
		name:              "TogetherAI",
		config:            config,
		TogetherAIHeaders: headers,
	}
}

func (togetherAI togetherAIProvider) Name() string {
	return togetherAI.name
}

func (togetherAI togetherAIProvider) Schema() []byte {
	return []byte{}
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
