package anyscale

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.ProviderInterface = &anyscaleProvider{}

type anyscaleProvider struct {
	Name   string
	Config base.ProviderConfig
	AnyscaleHeaders
}

func NewAnyscaleProvider(headers AnyscaleHeaders) *anyscaleProvider {
	config := getAnyscaleConfig("https://api.endpoints.anyscale.com")

	return &anyscaleProvider{
		Name:            "Anyscale",
		Config:          config,
		AnyscaleHeaders: headers,
	}
}

func (anyscale anyscaleProvider) GetName() string {
	return anyscale.Name
}

func (anyscale anyscaleProvider) Validate() error {
	return utils.ValidateHeaders(anyscale.AnyscaleHeaders)
}

func getAnyscaleConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
