package anyscale

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &anyscaleProvider{}

type anyscaleProvider struct {
	name   string
	config base.ProviderConfig
	conn   models.Connection
}

func (anyscale anyscaleProvider) Name() string {
	return anyscale.name
}

func (anyscale anyscaleProvider) Schema() []byte {
	return schema
}

func getAnyscaleConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	models.ProviderRegistry["anyscale"] = func(connection models.Connection) base.IProvider {
		config := getAnyscaleConfig("https://api.endpoints.anyscale.com")
		return &anyscaleProvider{
			name:   "Anyscale",
			config: config,
			conn:   connection,
		}
	}
}
