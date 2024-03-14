package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/requester"
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

func (oai *openAIProvider) ChatCompletions(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	client := requester.NewHTTPClient()

	payload.Stream = false
	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", oai.config.BaseURL, oai.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	req = oai.AddDefaultHeaders(req, provider.AuthorizationHeader)
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
	providerConfigMap := oai.provider.GetConfig([]string{
		provider.AuthorizationHeader,
	})

	var authorizationHeader string
	if val, ok := providerConfigMap[provider.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)
	return req
}

func (oai *openAIProvider) StreamChatCompletions(ctx context.Context, payload *chat.ChatCompletionRequest, stream chan []byte) error {
	client := requester.NewHTTPClient()

	payload.Stream = true
	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", oai.config.BaseURL, oai.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return err
	}

	req = oai.AddDefaultHeaders(req, provider.AuthorizationHeader)
	resp, err := client.SendRequestRaw(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		stream <- scanner.Bytes()

		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			event := strings.TrimPrefix(line, "data:")
			event = strings.TrimSpace(strings.Trim(event, "\n"))
			if strings.Contains(line, "[DONE]") {
				break
			}

			var data chat.ChatCompletionResponse
			if err := json.Unmarshal([]byte(event), &data); err != nil {
				return err
			}
		}
	}

	close(stream)
	return nil
}
