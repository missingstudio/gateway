package deepinfra

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &deepinfraProvider{}

type deepinfraProvider struct {
	name   string
	config base.ProviderConfig
	conn   models.Connection
}

func (deepinfra deepinfraProvider) Name() string {
	return deepinfra.name
}

func (deepinfra deepinfraProvider) Schema() []byte {
	return schema
}

func getDeepinfraConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/chat/completions",
	}
}

func init() {
	models.ProviderRegistry["deepinfra"] = func(connection models.Connection) base.IProvider {
		config := getDeepinfraConfig("https://api.deepinfra.com/v1/openai")
		return &deepinfraProvider{
			name:   "Deepinfra",
			config: config,
			conn:   connection,
		}
	}
}
