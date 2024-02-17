package togetherai

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

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
	return schema
}

func getTogetherAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
