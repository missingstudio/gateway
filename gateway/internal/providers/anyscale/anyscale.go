package anyscale

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/pkg/requester"
)

func (anyscale *anyscaleProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	client := requester.NewHTTPClient()

	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", anyscale.config.BaseURL, anyscale.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	req = anyscale.AddDefaultHeaders(req, connection.AuthorizationHeader)
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

func (anyscale *anyscaleProvider) AddDefaultHeaders(req *http.Request, key string) *http.Request {
	connectionConfigMap := anyscale.conn.GetConfig([]string{
		connection.AuthorizationHeader,
	})

	var authorizationHeader string
	if val, ok := connectionConfigMap[connection.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)
	return req
}

func (anyscale *anyscaleProvider) Models() []string {
	return []string{
		"meta-llama/Llama-2-7b-chat-hf",
		"meta-llama/Llama-2-13b-chat-hf",
		"meta-llama/Llama-2-70b-chat-hf",
		"codellama/CodeLlama-70b-Instruct-hf",
		"mistralai/Mistral-7B-Instruct-v0.1",
		"mistralai/Mixtral-8x7B-Instruct-v0.1",
		"mlabonne/NeuralHermes-2.5-Mistral-7B",
		"BAAI/bge-large-en-v1.5",
		"thenlper/gte-large",
	}
}
