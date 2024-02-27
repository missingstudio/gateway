package azure

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &azureProvider{}

type azureProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
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

func GetAzureInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title:       "Azure",
		Name:        "azure",
		Description: "Azure OpenAI Service offers industry-leading coding and language AI models that you can fine-tune to your specific needs for a variety of use cases.",
	}
}

func GetAzureConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}

func init() {}
