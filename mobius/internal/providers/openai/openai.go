package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

type OpenAIProviderFactory struct{}

func (f OpenAIProviderFactory) Create(token string) base.LLMProvider {
	openAIProvider := NewOpenAIProvider(token, "https://api.openai.com")
	return openAIProvider
}

type OpenAIProvider struct {
	APIKey string
	Config base.ProviderConfig
}

func NewOpenAIProvider(token string, baseURL string) *OpenAIProvider {
	config := getOpenAIConfig(baseURL)
	return &OpenAIProvider{
		APIKey: token,
		Config: config,
	}
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func (oai OpenAIProvider) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	payload, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "POST", oai.Config.BaseURL+oai.Config.ChatCompletions, bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oai.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data llmv1.CompletionResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
