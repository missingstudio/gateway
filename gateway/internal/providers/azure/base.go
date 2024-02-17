package azure

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

var _ base.IProvider = &azureProvider{}

type azureProvider struct {
	name   string
	config base.ProviderConfig
	AzureHeaders
}

func NewAzureProvider(headers AzureHeaders) *azureProvider {
	config := getAzureConfig()

	return &azureProvider{
		name:         "Azure",
		config:       config,
		AzureHeaders: headers,
	}
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
