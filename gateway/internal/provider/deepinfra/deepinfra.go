package deepinfra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/requester"
)

func (deepinfra *deepinfraProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	client := requester.NewHTTPClient()

	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", deepinfra.config.BaseURL, deepinfra.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	req = deepinfra.AddDefaultHeaders(req, provider.AuthorizationHeader)
	resp, err := client.SendRequestRaw(req)
	if err != nil {
		return nil, err
	}

	data := &chat.ChatCompletionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (deepinfra *deepinfraProvider) AddDefaultHeaders(req *http.Request, key string) *http.Request {
	providerConfigMap := deepinfra.provider.GetConfig([]string{key})

	var authorizationHeader string
	if val, ok := providerConfigMap[provider.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)
	return req
}

func (deepinfra *deepinfraProvider) Models() []string {
	return []string{
		"DeepInfra/pygmalion-13b-4bit-128g",
		"codellama/CodeLlama-70b-Instruct-hf",
		"cognitivecomputations/dolphin-2.6-mixtral-8x7b",
		"mistralai/Mixtral-8x7B-Instruct-v0.1",
		"lizpreciatior/lzlv_70b_fp16_hf",
		"deepinfra/airoboros-70b",
		"meta-llama/Llama-2-13b-chat-hf",
		"mistralai/Mistral-7B-Instruct-v0.1",
		"codellama/CodeLlama-34b-Instruct-hf",
		"meta-llama/Llama-2-70b-chat-hf",
		"meta-llama/Llama-2-7b-chat-hf",
		"jondurbin/airoboros-l2-70b-gpt4-1.4.1",
		"01-ai/Yi-34B-Chat",
		"Austism/chronos-hermes-13b-v2",
	}
}
