package togetherai

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &togetherAIProvider{}

type togetherAIProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   models.Connection
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
	models.ProviderRegistry["togetherai"] = func(connection models.Connection) base.IProvider {
		config := getTogetherAIConfig("https://api.together.xyz")
		return &togetherAIProvider{
			info:   getTogetherAIInfo(),
			config: config,
			conn:   connection,
		}
	}
}
