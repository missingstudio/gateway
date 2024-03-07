package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/missingstudio/ai/common/errors"
	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/pkg/utils"
	llmv1 "github.com/missingstudio/ai/protos/pkg/llm/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *V1Handler) ListAPIKeys(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[llmv1.ListAPIKeysResponse], error) {
	keys, err := s.apikeyService.GetAll(ctx)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	data := []*llmv1.APIKey{}
	for _, k := range keys {
		data = append(data, &llmv1.APIKey{
			Id:          k.Id.String(),
			Name:        k.Name,
			MaskedValue: utils.MaskString(k.Value),
			CreatedAt:   timestamppb.New(k.CreatedAt),
			LastUsedAt:  timestamppb.New(k.LastUsedAt),
		})
	}

	return connect.NewResponse(&llmv1.ListAPIKeysResponse{
		Keys: data,
	}), nil
}

func (s *V1Handler) CreateAPIKey(ctx context.Context, req *connect.Request[llmv1.CreateAPIKeyRequest]) (*connect.Response[llmv1.CreateAPIKeyResponse], error) {
	securekey, err := utils.GenerateSecureAPIKey()
	if err != nil {
		return nil, errors.New(err)
	}

	key := apikey.APIKey{
		Name:  req.Msg.Name,
		Value: securekey,
	}

	newkey, err := s.apikeyService.Create(ctx, key)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&llmv1.CreateAPIKeyResponse{
		Key: &llmv1.APIKey{
			Id:         newkey.Id.String(),
			Name:       newkey.Name,
			Value:      newkey.Value,
			CreatedAt:  timestamppb.New(newkey.CreatedAt),
			LastUsedAt: timestamppb.New(newkey.LastUsedAt),
		},
	}), nil
}

func (s *V1Handler) GetAPIKey(ctx context.Context, req *connect.Request[llmv1.GetAPIKeyRequest]) (*connect.Response[llmv1.GetAPIKeyResponse], error) {
	key, err := s.apikeyService.Get(ctx, req.Msg.Id)
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	k := &llmv1.APIKey{
		Id:          key.Id.String(),
		Name:        key.Name,
		MaskedValue: utils.MaskString(key.Value),
		CreatedAt:   timestamppb.New(key.CreatedAt),
		LastUsedAt:  timestamppb.New(key.LastUsedAt),
	}

	return connect.NewResponse(&llmv1.GetAPIKeyResponse{
		Key: k,
	}), nil
}

func (s *V1Handler) UpdateAPIKey(ctx context.Context, req *connect.Request[llmv1.UpdateAPIKeyRequest]) (*connect.Response[llmv1.UpdateAPIKeyResponse], error) {
	parsedUUID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, errors.New(err)
	}

	key := apikey.APIKey{
		Id:   parsedUUID,
		Name: req.Msg.Name,
	}

	updatedkey, err := s.apikeyService.Update(ctx, key)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&llmv1.UpdateAPIKeyResponse{
		Key: &llmv1.APIKey{
			Id:          updatedkey.Id.String(),
			Name:        updatedkey.Name,
			MaskedValue: utils.MaskString(updatedkey.Value),
			CreatedAt:   timestamppb.New(updatedkey.CreatedAt),
			LastUsedAt:  timestamppb.New(updatedkey.LastUsedAt),
		},
	}), nil
}

func (s *V1Handler) DeleteAPIKey(ctx context.Context, req *connect.Request[llmv1.DeleteAPIKeyRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.apikeyService.DeleteByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
