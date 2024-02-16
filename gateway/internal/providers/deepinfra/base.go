package deepinfra

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.ProviderInterface = &deepinfraProvider{}

type deepinfraProvider struct {
	Name   string
	Config base.ProviderConfig
	DeepinfraHeaders
}

func NewDeepinfraProvider(headers DeepinfraHeaders) *deepinfraProvider {
	config := getDeepinfraConfig("https://api.deepinfra.com/v1/openai")

	return &deepinfraProvider{
		Name:             "Deepinfra",
		Config:           config,
		DeepinfraHeaders: headers,
	}
}

func (deepinfra deepinfraProvider) GetName() string {
	return deepinfra.Name
}

func (deepinfra deepinfraProvider) Validate() error {
	return utils.ValidateHeaders(deepinfra.DeepinfraHeaders)
}

func getDeepinfraConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/chat/completions",
	}
}
