package provider

import (
	"errors"

	"github.com/missingstudio/ai/gateway/core/provider"
	_ "github.com/missingstudio/ai/gateway/internal/provider/anyscale"
	_ "github.com/missingstudio/ai/gateway/internal/provider/azure"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
	_ "github.com/missingstudio/ai/gateway/internal/provider/deepinfra"
	_ "github.com/missingstudio/ai/gateway/internal/provider/groq"
	_ "github.com/missingstudio/ai/gateway/internal/provider/openai"
	_ "github.com/missingstudio/ai/gateway/internal/provider/togetherai"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GetProviders returns all supported providers
func (s Service) GetProviders() map[string]base.Provider {
	providers := map[string]base.Provider{}

	for name, p := range base.ProviderRegistry {
		providers[name] = p(provider.Provider{})
	}
	return providers
}

func (s Service) GetProvider(provider provider.Provider) (base.Provider, error) {
	if val, ok := base.ProviderRegistry[provider.Name]; ok {
		return val(provider), nil
	}
	return nil, errors.New("unsupported provider")
}
