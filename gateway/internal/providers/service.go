package providers

import (
	"errors"

	_ "github.com/missingstudio/studio/backend/internal/providers/anyscale"
	_ "github.com/missingstudio/studio/backend/internal/providers/azure"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	_ "github.com/missingstudio/studio/backend/internal/providers/deepinfra"
	_ "github.com/missingstudio/studio/backend/internal/providers/openai"
	_ "github.com/missingstudio/studio/backend/internal/providers/togetherai"
	"github.com/missingstudio/studio/backend/models"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GetProviders returns all supported providers
func (s Service) GetProviders() map[string]base.IProvider {
	providers := map[string]base.IProvider{}

	for name, p := range models.ProviderRegistry {
		providers[name] = p(models.Connection{})
	}
	return providers
}

func (s Service) GetProvider(conn models.Connection) (base.IProvider, error) {
	if val, ok := models.ProviderRegistry[conn.Name]; ok {
		return val(conn), nil
	}
	return nil, errors.New("unsupported connection")
}
