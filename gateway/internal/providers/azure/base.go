package azure

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &azureProvider{}

type azureProvider struct {
	name   string
	config base.ProviderConfig
	conn   models.Connection
}

func (az azureProvider) Name() string {
	return az.name
}

func (az azureProvider) Schema() []byte {
	return schema
}

func getAzureConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}

func init() {
	models.ProviderRegistry["azure"] = func(connection models.Connection) base.IProvider {
		config := getAzureConfig()
		return &azureProvider{
			name:   "Azure",
			config: config,
			conn:   connection,
		}
	}
}
