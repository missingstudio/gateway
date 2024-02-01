package azure

import (
	"net/http"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type AzureProviderFactory struct{}

type AzureHeaders struct {
	APIKey       string `validate:"required" json:"Authorization" error:"API key is required"`
	ResourceName string `validate:"required" json:"X-Ms-Azure-Resource-Name" error:"Resource Name is required"`
	DeploymentID string `validate:"required" json:"X-Ms-Deployment-ID" error:"Deployment ID is required"`
	APIVersion   string `validate:"required" json:"X-Ms-API-Version" error:"API Version is required"`
}

func (azf AzureProviderFactory) Validate(headers http.Header) (*AzureHeaders, error) {
	var azHeaders AzureHeaders
	if err := utils.UnmarshalHeader(headers, &azHeaders); err != nil {
		return nil, errors.New(err)
	}

	if err := utils.ValidateHeaders(azHeaders); err != nil {
		return nil, err
	}

	return &azHeaders, nil
}

func (azf AzureProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	azureHeaders, err := azf.Validate(headers)
	if err != nil {
		return nil, err
	}

	azureProvider := NewAzureProvider(*azureHeaders)
	return azureProvider, nil
}

type AzureProvider struct {
	Name string
	AzureHeaders
	Config base.ProviderConfig
}

func NewAzureProvider(headers AzureHeaders) *AzureProvider {
	config := getAzureConfig()

	return &AzureProvider{
		Name:         "Azure",
		Config:       config,
		AzureHeaders: headers,
	}
}

func (az *AzureProvider) GetName() string {
	return az.Name
}

func getAzureConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}
