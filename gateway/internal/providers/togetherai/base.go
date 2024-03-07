package togetherai

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/internal/providers/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionInterface = &togetherAIProvider{}

type togetherAIProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   connection.Connection
}

func (anyscale togetherAIProvider) Info() base.ProviderInfo {
	return anyscale.info
}

func (togetherAI togetherAIProvider) Config() base.ProviderConfig {
	return togetherAI.config
}

func (togetherAI togetherAIProvider) Schema() []byte {
	return schema
}

func getTogetherAIInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title:       "Together AI",
		Name:        "togetherai",
		Description: `Build gen AI models with Together AI. Benefit from the fastest and most cost-efficient tools and infra.`,
	}
}

func getTogetherAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["togetherai"] = func(connection connection.Connection) base.IProvider {
		config := getTogetherAIConfig("https://api.together.xyz")
		return &togetherAIProvider{
			info:   getTogetherAIInfo(),
			config: config,
			conn:   connection,
		}
	}
}
