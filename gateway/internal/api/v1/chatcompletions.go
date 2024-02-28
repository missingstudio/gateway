package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/router"
	"github.com/missingstudio/studio/backend/models"
	"github.com/missingstudio/studio/common/errors"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

var (
	ErrChatCompletionStreamNotSupported = errors.NewBadRequest("streaming is not supported with this method, please use StreamChatCompletions")
	ErrChatCompletionNotSupported       = errors.NewInternalError("provider don't have chat Completion capabilities")
	ErrNoProviderAbleToServe            = errors.NewInternalError("none of the provider able to serve you")
	ErrRequiredHeaderNotExit            = errors.NewBadRequest(fmt.Sprintf("%s header is required", constants.XMSConfig))
)

func (s *V1Handler) ChatCompletions(
	ctx context.Context,
	req *connect.Request[llmv1.ChatCompletionRequest],
) (*connect.Response[llmv1.ChatCompletionResponse], error) {
	// Check if required headers are available
	routerConfig := router.GetContextWithRouterConfig(ctx)
	if routerConfig == nil {
		return nil, ErrRequiredHeaderNotExit
	}

	chatCompletionRequestSchema, err := s.createChatCompletionRequestSchema(req.Msg)
	if err != nil {
		return nil, errors.New(err)
	}

	startTime := time.Now()
	rsvc := router.NewRouterService(routerConfig)

	providerConfig := rsvc.Next()
	if providerConfig == nil {
		return nil, ErrNoProviderAbleToServe
	}

	authConfig := map[string]any{"auth": providerConfig.Auth}
	connectionObj := models.Connection{
		Name:   providerConfig.Name,
		Config: authConfig,
	}

	p, err := s.providerService.GetProvider(connectionObj)
	if err != nil {
		return nil, errors.New(err)
	}

	// Validate provider configs
	err = providers.Validate(p, authConfig)
	if err != nil {
		return nil, errors.NewBadRequest(err.Error())
	}

	chatCompletionProvider, ok := p.(base.ChatCompletionInterface)
	if !ok {
		return nil, ErrChatCompletionNotSupported
	}

	resp, err := chatCompletionProvider.ChatCompletion(ctx, chatCompletionRequestSchema)
	if err != nil {
		return nil, errors.New(err)
	}

	providerInfo := p.Info()
	latency := time.Since(startTime)
	s.sendMetrics(providerInfo.Name, latency, resp)

	chatCompletionResponseSchema, err := s.createChatCompletionResponseSchema(resp)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(chatCompletionResponseSchema), nil
}

func (s *V1Handler) createChatCompletionRequestSchema(req *llmv1.ChatCompletionRequest) (*models.ChatCompletionRequest, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data := &models.ChatCompletionRequest{}
	err = json.Unmarshal(payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *V1Handler) createChatCompletionResponseSchema(resp *models.ChatCompletionResponse) (*llmv1.ChatCompletionResponse, error) {
	payload, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	data := &llmv1.ChatCompletionResponse{}
	err = json.Unmarshal(payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *V1Handler) sendMetrics(provider string, latency time.Duration, response *models.ChatCompletionResponse) {
	ingesterdata := make(map[string]any)
	ingesterdata["provider"] = provider
	ingesterdata["latency"] = latency
	ingesterdata["model"] = response.Model
	ingesterdata["total_tokens"] = response.Usage.TotalTokens
	ingesterdata["prompt_tokens"] = response.Usage.PromptTokens
	ingesterdata["completion_tokens"] = response.Usage.CompletionTokens
	go s.ingester.Ingest(ingesterdata, "analytics")
}
