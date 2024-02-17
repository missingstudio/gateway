package azure

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

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

func (togetherAI azureProvider) Schema() []byte {
	return []byte{}
}

func (az azureProvider) Validate() error {
	return utils.ValidateHeaders(az.AzureHeaders)
}

func getAzureConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}
