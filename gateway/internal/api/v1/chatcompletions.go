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

	startTime := time.Now()
	payload, err := json.Marshal(req.Msg)
	if err != nil {
		return nil, errors.New(err)
	}

	rsvc := router.NewRouterService(routerConfig)

	data := &llmv1.ChatCompletionResponse{}
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

	resp, err := chatCompletionProvider.ChatCompletion(ctx, payload)
	if err != nil {
		return nil, errors.New(err)
	}

	latency := time.Since(startTime)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, errors.New(err)
	}

	ingesterdata := make(map[string]any)
	providerInfo := p.Info()

	ingesterdata["provider"] = providerInfo.Name
	ingesterdata["model"] = data.Model
	ingesterdata["latency"] = latency
	ingesterdata["total_tokens"] = *data.Usage.TotalTokens
	ingesterdata["prompt_tokens"] = *data.Usage.PromptTokens
	ingesterdata["completion_tokens"] = *data.Usage.CompletionTokens

	go s.ingester.Ingest(ingesterdata, "analytics")

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(data), nil
}
