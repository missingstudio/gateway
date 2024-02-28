package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (s *V1Handler) ListModels(ctx context.Context, req *connect.Request[llmv1.ModelRequest]) (*connect.Response[llmv1.ModelResponse], error) {
	allProviderModels := map[string]*llmv1.ProviderModels{}

	for name := range base.ProviderRegistry {
		provider, err := s.providerService.GetProvider(models.Connection{Name: name})
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
