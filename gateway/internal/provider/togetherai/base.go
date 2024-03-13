package togetherai

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &togetherAIProvider{}

type togetherAIProvider struct {
	info     base.Info
	config   base.Config
	provider provider.Provider
}

func (anyscale togetherAIProvider) Info() base.Info {
	return anyscale.info
}

func (togetherAI togetherAIProvider) Config() base.Config {
	return togetherAI.config
}

func (togetherAI togetherAIProvider) Schema() []byte {
	return schema
}

func getTogetherAIInfo() base.Info {
	return base.Info{
		Title:       "Together AI",
		Name:        "togetherai",
		Description: `Build gen AI models with Together AI. Benefit from the fastest and most cost-efficient tools and infra.`,
	}
}

func getTogetherAIConfig(baseURL string) base.Config {
	return base.Config{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["togetherai"] = func(provider provider.Provider) base.Provider {
		config := getTogetherAIConfig("https://api.together.xyz")
		return &togetherAIProvider{
			info:     getTogetherAIInfo(),
			config:   config,
			provider: provider,
		}
	}
}
