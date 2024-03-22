package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
	llmv1 "github.com/missingstudio/protos/pkg/llm/v1"
)

func (s *V1Handler) ListModels(ctx context.Context, req *connect.Request[llmv1.ModelRequest]) (*connect.Response[llmv1.ModelResponse], error) {
	allProviderModels := map[string]*llmv1.ProviderModels{}

	for name := range base.ProviderRegistry {
		// Check if the provider is healthy before fetching models
		if !router.DefaultHealthChecker{}.IsHealthy(name) {
			continue
		}

		provider, err := s.iProviderService.GetProvider(provider.Provider{Name: name})
		if err != nil {
			continue
		}

		providerInfo := provider.Info()
		providerModels := provider.Models()
		providerName := providerInfo.Name

		var models []*llmv1.Model
		for _, val := range providerModels {
			models = append(models, &llmv1.Model{
				Name:  val,
				Value: val,
			})
		}

		allProviderModels[name] = &llmv1.ProviderModels{
			Name:   providerName,
			Models: models,
		}
	}

	return connect.NewResponse(&llmv1.ModelResponse{
		Models: allProviderModels,
	}), nil
}
