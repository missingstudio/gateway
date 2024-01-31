package azure

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type AzureProviderFactory struct{}

type AzureHeaders struct {
	APIKey       string `validate:"required" json:"Authorization"`
	ResourceName string `validate:"required" json:"X-Ms-Azure-Resource-Name"`
	DeploymentID string `validate:"required" json:"X-Ms-Deployment-ID"`
	APIVersion   string `validate:"required" json:"X-Ms-API-Version"`
}

func (azf AzureProviderFactory) Validate(headers http.Header) (*AzureHeaders, error) {
	var azHeaders AzureHeaders
	if err := utils.UnmarshalHeader(headers, &azHeaders); err != nil {
		return nil, errors.New(err)
	}

	validate := validator.New()
	if err := validate.Struct(azHeaders); err != nil {
		return nil, errors.NewBadRequest("provider's required headers are missing")
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
		Name:         "Azure AI",
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
