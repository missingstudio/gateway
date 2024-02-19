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
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   models.Connection
}

func (anyscale azureProvider) Info() base.ProviderInfo {
	return anyscale.info
}

func (az azureProvider) Config() base.ProviderConfig {
	return az.config
}

func (az azureProvider) Schema() []byte {
	return schema
}

func getAzureInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title:       "Azure",
		Name:        "azure",
		Description: "Azure OpenAI Service offers industry-leading coding and language AI models that you can fine-tune to your specific needs for a variety of use cases.",
	}
}

func getAzureConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}

func init() {}
