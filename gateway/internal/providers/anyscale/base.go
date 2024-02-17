package anyscale

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.IProvider = &anyscaleProvider{}

type anyscaleProvider struct {
	name   string
	config base.ProviderConfig
	AnyscaleHeaders
}

func NewAnyscaleProvider(headers AnyscaleHeaders) *anyscaleProvider {
	config := getAnyscaleConfig("https://api.endpoints.anyscale.com")

	return &anyscaleProvider{
		name:            "Anyscale",
		config:          config,
		AnyscaleHeaders: headers,
	}
}

func (anyscale anyscaleProvider) Name() string {
	return anyscale.name
}

func (togetherAI anyscaleProvider) Schema() []byte {
	return []byte{}
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
