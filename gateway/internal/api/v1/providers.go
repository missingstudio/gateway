package v1

import (
	"context"
	"encoding/json"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/models"
	"github.com/missingstudio/studio/common/errors"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *V1Handler) ListProviders(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[llmv1.ProvidersResponse], error) {
	providers := s.providerService.GetProviders()

	data := []*llmv1.Provider{}
	for _, provider := range providers {
		providerInfo := provider.Info()
		data = append(data, &llmv1.Provider{
			Title:       providerInfo.Title,
			Name:        providerInfo.Name,
			Description: providerInfo.Description,
		})
	}

	return connect.NewResponse(&llmv1.ProvidersResponse{
		Providers: data,
	}), nil
}

func (s *V1Handler) GetProviderById(ctx context.Context, req *connect.Request[llmv1.GetProviderRequest]) (*connect.Response[llmv1.GetProviderResponse], error) {
	provider, err := s.providerService.GetProvider(models.Connection{Name: req.Msg.Id})
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	info := provider.Info()
	p := &llmv1.Provider{
		Title:       info.Title,
		Name:        info.Name,
		Description: info.Description,
	}

	return connect.NewResponse(&llmv1.GetProviderResponse{
		Provider: p,
	}), nil
}

func (s *V1Handler) GetProviderConfig(ctx context.Context, req *connect.Request[llmv1.GetProviderConfigRequest]) (*connect.Response[llmv1.GetProviderConfigResponse], error) {
	provider, err := s.providerService.GetProvider(models.Connection{Name: req.Msg.Id})
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	configs := map[string]any{}
	if err := json.Unmarshal(provider.Schema(), &configs); err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	stConfigs, err := structpb.NewStruct(configs)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return connect.NewResponse(&llmv1.GetProviderConfigResponse{
		Config: stConfigs,
	}), nil
}
