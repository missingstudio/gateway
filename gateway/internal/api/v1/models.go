package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (s *V1Handler) ListModels(ctx context.Context, req *connect.Request[llmv1.ModelRequest]) (*connect.Response[llmv1.ModelResponse], error) {
	plist := providers.Providers
	allProviderModels := map[string]*llmv1.ProviderModels{}

	for _, p := range plist {
		providerfactory, err := providers.NewProvider(p, req.Header())
		if err != nil {
			continue
		}

		providerName := providerfactory.Name()
		providerModels := providerfactory.Models()

		var models []*llmv1.Model
		for _, val := range providerModels {
			models = append(models, &llmv1.Model{
				Name:  val,
				Value: val,
			})
		}

		allProviderModels[p] = &llmv1.ProviderModels{
			Name:   providerName,
			Models: models,
		}
	}

	return connect.NewResponse(&llmv1.ModelResponse{
		Models: allProviderModels,
	}), nil
}
