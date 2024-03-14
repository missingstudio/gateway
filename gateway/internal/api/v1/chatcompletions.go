package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/constants"
	iProvider "github.com/missingstudio/ai/gateway/internal/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
	"github.com/missingstudio/ai/gateway/internal/router"
	"github.com/missingstudio/common/errors"
	llmv1 "github.com/missingstudio/protos/pkg/llm/v1"
	"google.golang.org/genproto/googleapis/api/httpbody"
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
	connectionObj := provider.Provider{
		Name:   providerConfig.Name,
		Config: authConfig,
	}

	p, err := s.iProviderService.GetProvider(connectionObj)
	if err != nil {
		return nil, errors.New(err)
	}

	// Validate provider configs
	err = iProvider.Validate(p, authConfig)
	if err != nil {
		return nil, errors.NewBadRequest(err.Error())
	}

	chatCompletionProvider, ok := p.(base.ChatCompletionProvider)
	if !ok {
		return nil, ErrChatCompletionNotSupported
	}

	resp, err := chatCompletionProvider.ChatCompletions(ctx, chatCompletionRequestSchema)
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

func (s *V1Handler) StreamChatCompletions(ctx context.Context, req *connect.Request[llmv1.ChatCompletionRequest], stream *connect.ServerStream[httpbody.HttpBody]) error {
	// Check if required headers are available
	routerConfig := router.GetContextWithRouterConfig(ctx)
	if routerConfig == nil {
		return ErrRequiredHeaderNotExit
	}

	chatCompletionRequestSchema, err := s.createChatCompletionRequestSchema(req.Msg)
	if err != nil {
		return errors.New(err)
	}

	rsvc := router.NewRouterService(routerConfig)

	providerConfig := rsvc.Next()
	if providerConfig == nil {
		return ErrNoProviderAbleToServe
	}

	authConfig := map[string]any{"auth": providerConfig.Auth}
	connectionObj := provider.Provider{
		Name:   providerConfig.Name,
		Config: authConfig,
	}

	p, err := s.iProviderService.GetProvider(connectionObj)
	if err != nil {
		return errors.New(err)
	}

	// Validate provider configs
	err = iProvider.Validate(p, authConfig)
	if err != nil {
		return errors.NewBadRequest(err.Error())
	}

	chatCompletionProvider, ok := p.(base.StreamChatCompletionsProvider)
	if !ok {
		return ErrChatCompletionNotSupported
	}

	byteChannel := make(chan []byte)
	go sendDataStream(byteChannel, stream)
	return chatCompletionProvider.StreamChatCompletions(ctx, chatCompletionRequestSchema, byteChannel)
}

func sendDataStream(byteChannel chan []byte, stream *connect.ServerStream[httpbody.HttpBody]) {
	for data := range byteChannel {
		_ = stream.Send(&httpbody.HttpBody{
			Data: data,
		})
	}
}

func (s *V1Handler) createChatCompletionRequestSchema(req *llmv1.ChatCompletionRequest) (*chat.ChatCompletionRequest, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data := &chat.ChatCompletionRequest{}
	err = json.Unmarshal(payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *V1Handler) createChatCompletionResponseSchema(resp *chat.ChatCompletionResponse) (*llmv1.ChatCompletionResponse, error) {
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

func (s *V1Handler) sendMetrics(provider string, latency time.Duration, response *chat.ChatCompletionResponse) {
	ingesterdata := make(map[string]any)
	ingesterdata["provider"] = provider
	ingesterdata["latency"] = latency
	ingesterdata["model"] = response.Model
	ingesterdata["total_tokens"] = response.Usage.TotalTokens
	ingesterdata["prompt_tokens"] = response.Usage.PromptTokens
	ingesterdata["completion_tokens"] = response.Usage.CompletionTokens
	go s.ingester.Ingest(ingesterdata)
}
