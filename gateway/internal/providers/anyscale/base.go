package anyscale

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

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

func (anyscale anyscaleProvider) Schema() []byte {
	return schema
}

func getAnyscaleConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
