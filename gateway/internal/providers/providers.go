package providers

import (
	"context"
	"net/http"

	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
	"github.com/missingstudio/studio/backend/internal/providers/anyscale"
	"github.com/missingstudio/studio/backend/internal/providers/azure"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/providers/deepinfra"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
	"github.com/missingstudio/studio/backend/internal/providers/togetherai"
)

type ProviderFactory interface {
	Create(headers http.Header) (base.ProviderInterface, error)
}

var ProviderFactories = make(map[string]ProviderFactory)

func init() {
	ProviderFactories["openai"] = openai.OpenAIProviderFactory{}
	ProviderFactories["azure"] = azure.AzureProviderFactory{}
	ProviderFactories["anyscale"] = anyscale.AnyscaleProviderFactory{}
	ProviderFactories["deepinfra"] = deepinfra.DeepinfraProviderFactory{}
	ProviderFactories["togetherai"] = togetherai.TogetherAIProviderFactory{}
}

func GetProvider(ctx context.Context, headers http.Header) (base.ProviderInterface, error) {
	providerName := headers.Get(constants.XMSProvider)
	if providerName == "" {
		return nil, errors.ErrProviderHeaderNotExit
	}

	providerFactory, ok := ProviderFactories[providerName]
	if !ok {
		return nil, errors.ErrProviderNotFound
	}

	return providerFactory.Create(headers)
}
