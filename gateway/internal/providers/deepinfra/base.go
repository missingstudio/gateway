package deepinfra

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

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

func (deepinfra deepinfraProvider) Schema() []byte {
	return schema
}

func getDeepinfraConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/chat/completions",
	}
}
