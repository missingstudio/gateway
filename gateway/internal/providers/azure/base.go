package azure

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.ProviderInterface = &azureProvider{}

type azureProvider struct {
	Name string
	AzureHeaders
	Config base.ProviderConfig
}

func NewAzureProvider(headers AzureHeaders) *azureProvider {
	config := getAzureConfig()

	return &azureProvider{
		Name:         "Azure",
		Config:       config,
		AzureHeaders: headers,
	}
}

func (az azureProvider) GetName() string {
	return az.Name
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
