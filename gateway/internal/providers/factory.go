package providers

import (
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/internal/providers/anyscale"
	"github.com/missingstudio/studio/backend/internal/providers/azure"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/providers/deepinfra"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
	"github.com/missingstudio/studio/backend/internal/providers/togetherai"
)

const (
	Openai     = "openai"
	Azure      = "azure"
	Anyscale   = "anyscale"
	Deepinfra  = "deepinfra"
	Togetherai = "togetherai"
)

var Providers = []string{Openai, Azure, Anyscale, Deepinfra, Togetherai}

// NewProvider initializes the provider instance based on Config
func NewProvider(provider string, headers http.Header) (base.ProviderInterface, error) {
	switch provider {
	case Openai:
		providerFactory := openai.OpenAIProviderFactory{}
		return providerFactory.Create(headers)
	case Azure:
		providerFactory := azure.AzureProviderFactory{}
		return providerFactory.Create(headers)
	case Anyscale:
		providerFactory := anyscale.AnyscaleProviderFactory{}
		return providerFactory.Create(headers)
	case Togetherai:
		providerFactory := togetherai.TogetherAIProviderFactory{}
		return providerFactory.Create(headers)
	case Deepinfra:
		providerFactory := deepinfra.DeepinfraProviderFactory{}
		return providerFactory.Create(headers)
	default:
		return nil, fmt.Errorf("Unknown provider: %s", provider)
	}
}
