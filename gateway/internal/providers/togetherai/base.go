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
	name   string
	config base.ProviderConfig
	conn   models.Connection
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

func init() {
	models.ProviderRegistry["togetherai"] = func(connection models.Connection) base.IProvider {
		config := getTogetherAIConfig("https://api.together.xyz")
		return &togetherAIProvider{
			name:   "Together AI",
			config: config,
			conn:   connection,
		}
	}
}
