package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/core/chat"
	"github.com/missingstudio/studio/backend/core/connection"
	"github.com/missingstudio/studio/backend/pkg/requester"
)

var OpenAIModels = []string{
	"gpt-4-0125-preview",
	"gpt-4-turbo-preview",
	"gpt-4-1106-preview",
	"gpt-4-vision-preview",
	"gpt-4-1106-vision-preview",
	"gpt-4",
	"gpt-4-0613",
	"gpt-4-32k",
	"gpt-4-32k-0613",
	"gpt-3.5-turbo-0125",
	"gpt-3.5-turbo",
	"gpt-3.5-turbo-1106",
	"gpt-3.5-turbo-instruct",
}

func (oai *openAIProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	client := requester.NewHTTPClient()

	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", oai.config.BaseURL, oai.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	req = oai.AddDefaultHeaders(req, connection.AuthorizationHeader)
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

func (*openAIProvider) Models() []string {
	return OpenAIModels
}

func (oai *openAIProvider) AddDefaultHeaders(req *http.Request, key string) *http.Request {
	connectionConfigMap := oai.conn.GetConfig([]string{key})

	var authorizationHeader string
	if val, ok := connectionConfigMap[connection.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)
	return req
}
