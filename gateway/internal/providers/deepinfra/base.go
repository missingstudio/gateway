package deepinfra

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

var _ base.IProvider = &deepinfraProvider{}

type deepinfraProvider struct {
	name   string
	config base.ProviderConfig
	DeepinfraHeaders
}

func NewDeepinfraProvider(headers DeepinfraHeaders) *deepinfraProvider {
	config := getDeepinfraConfig("https://api.deepinfra.com/v1/openai")

	return &deepinfraProvider{
		name:             "Deepinfra",
		config:           config,
		DeepinfraHeaders: headers,
	}
}

func (deepinfra deepinfraProvider) Name() string {
	return deepinfra.name
}

func (togetherAI deepinfraProvider) Schema() []byte {
	return []byte{}
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
