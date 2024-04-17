package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/huggingface"
	"github.com/missingstudio/common/errors"
	llmv1 "github.com/missingstudio/protos/pkg/llm/v1"
)

func (s *V1Handler) InitiateFineTuning(ctx context.Context, req *connect.Request[llmv1.FineTuneRequest]) (*connect.Response[llmv1.FineTuneResponse], error) {
	hfProvider, err := s.iProviderService.GetProvider(provider.Provider{Name: "HuggingFace"})
	if err != nil {
		return nil, errors.NewInternal("failed to get HuggingFace provider")
	}

	jobID, err := hfProvider.(*huggingface.HuggingFaceProvider).InitiateFineTuning(ctx, req.Payload.Model, req.Payload.Parameters)
	if err != nil {
		return nil, errors.NewInternal("failed to initiate fine-tuning: " + err.Error())
	}

	return connect.NewResponse(&llmv1.FineTuneResponse{
		JobId: jobID,
	}), nil
}

func (s *V1Handler) CheckFineTuningStatus(ctx context.Context, req *connect.Request[llmv1.FineTuneStatusRequest]) (*connect.Response[llmv1.FineTuneStatusResponse], error) {
	hfProvider, err := s.iProviderService.GetProvider(provider.Provider{Name: "HuggingFace"})
	if err != nil {
		return nil, errors.NewInternal("failed to get HuggingFace provider")
	}

	result, err := hfProvider.(*huggingface.HuggingFaceProvider).RetrieveFineTuningResults(ctx, req.Payload.JobId)
	if err != nil {
		return nil, errors.NewInternal("failed to retrieve fine-tuning results: " + err.Error())
	}

	status, ok := result["status"].(string)
	if !ok {
		return nil, errors.NewInternal("unexpected response format from HuggingFace")
	}

	return connect.NewResponse(&llmv1.FineTuneStatusResponse{
		Status: status,
	}), nil
}
