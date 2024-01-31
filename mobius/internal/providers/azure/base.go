package azure

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/common/errors"
)

type AzureProviderFactory struct{}

func (f AzureProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	authorization := headers.Get(config.Authorization)
	if authorization == "" {
		return nil, errors.NewBadRequest("authorization header is required")
	}

	authorizationKey := strings.Replace(authorization, "Bearer ", "", 1)
	azureProvider := NewazureProvider(authorizationKey)
	return azureProvider, nil
}

type AzureHeaders struct {
	APIKey string
}

type AzureProvider struct {
	Name   string
	Config base.ProviderConfig
	AzureHeaders
}

func NewazureProvider(apikey string) *AzureProvider {
	config := getAzureConfig()

	return &AzureProvider{
		Name: "Azure AI",
		AzureHeaders: AzureHeaders{
			APIKey: apikey,
		},
		Config: config,
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
