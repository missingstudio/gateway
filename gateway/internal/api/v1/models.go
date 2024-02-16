package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/common/errors"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *V1Handler) Models(ctx context.Context, req *connect.Request[llmv1.ModelRequest]) (*connect.Response[llmv1.ModelResponse], error) {
	plist := providers.Providers
	allProviderModels := map[string]interface{}{}

	for _, p := range plist {
		providerfactory, err := providers.NewProvider(p, req.Header())
		if err != nil {
			continue
		}

		providerModels := map[string]interface{}{
			"name":   providerfactory.GetName(),
			"models": providerfactory.GetModels(),
		}

		allProviderModels[p] = providerModels
	}

	m, err := structpb.NewStruct(allProviderModels)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&llmv1.ModelResponse{
		Models: m,
	}), nil
}
