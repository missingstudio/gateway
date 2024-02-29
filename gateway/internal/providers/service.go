package providers

import (
	"errors"

	"github.com/missingstudio/studio/backend/core/connection"
	_ "github.com/missingstudio/studio/backend/internal/providers/anyscale"
	_ "github.com/missingstudio/studio/backend/internal/providers/azure"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	_ "github.com/missingstudio/studio/backend/internal/providers/deepinfra"
	_ "github.com/missingstudio/studio/backend/internal/providers/openai"
	_ "github.com/missingstudio/studio/backend/internal/providers/togetherai"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GetProviders returns all supported providers
func (s Service) GetProviders() map[string]base.IProvider {
	providers := map[string]base.IProvider{}

	for name, p := range base.ProviderRegistry {
		providers[name] = p(connection.Connection{})
	}
	return providers
}

func (s Service) GetProvider(conn connection.Connection) (base.IProvider, error) {
	if val, ok := base.ProviderRegistry[conn.Name]; ok {
		return val(conn), nil
	}
	return nil, errors.New("unsupported connection")
}
