package openai

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &openAIProvider{}

type openAIProvider struct {
	name   string
	config base.ProviderConfig
	conn   models.Connection
}

func (oai openAIProvider) Name() string {
	return oai.name
}

func (oai openAIProvider) Schema() []byte {
	return schema
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	models.ProviderRegistry["openai"] = func(connection models.Connection) base.IProvider {
		config := getOpenAIConfig("https://api.openai.com")
		return &openAIProvider{
			name:   "OpenAI",
			config: config,
			conn:   connection,
		}
	}
}
