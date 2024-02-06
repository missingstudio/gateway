package providers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/providers/anyscale"
	"github.com/missingstudio/studio/backend/internal/providers/azure"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/providers/deepinfra"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
	"github.com/missingstudio/studio/backend/internal/providers/togetherai"
	"github.com/missingstudio/studio/common/errors"
)

var (
	ErrProviderHeaderNotExit = errors.New(fmt.Errorf("x-ms-provider provider header not available"))
	ErrProviderNotFound      = errors.NewNotFound("provider is not found")
)

type ProviderFactory interface {
	Create(headers http.Header) (base.ProviderInterface, error)
}

var providerFactories = make(map[string]ProviderFactory)

func init() {
	providerFactories["openai"] = openai.OpenAIProviderFactory{}
	providerFactories["azure"] = azure.AzureProviderFactory{}
	providerFactories["anyscale"] = anyscale.AnyscaleProviderFactory{}
	providerFactories["deepinfra"] = deepinfra.DeepinfraProviderFactory{}
	providerFactories["togetherai"] = togetherai.TogetherAIProviderFactory{}
}

func GetProvider(ctx context.Context, headers http.Header) (base.ProviderInterface, error) {
	providerName := headers.Get(config.XMSProvider)
	if providerName == "" {
		return nil, ErrProviderHeaderNotExit
	}

	providerFactory, ok := providerFactories[providerName]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return providerFactory.Create(headers)
}
