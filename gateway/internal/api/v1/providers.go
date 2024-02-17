package v1

import (
	"context"

	"connectrpc.com/connect"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *V1Handler) ListProviders(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[llmv1.ProvidersResponse], error) {
	providers := s.providerService.GetProviders()

	data := []*llmv1.Provider{}
	for name := range providers {
		data = append(data, &llmv1.Provider{
			Name: name,
		})
	}

	return connect.NewResponse(&llmv1.ProvidersResponse{
		Providers: data,
	}), nil
}
