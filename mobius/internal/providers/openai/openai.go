package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/missingstudio/studio/backend/pkg/requester"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (oai *OpenAIProvider) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	payload, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}

	client := requester.NewHTTPClient()
	req, _ := http.NewRequestWithContext(ctx, "POST", oai.Config.BaseURL+oai.Config.ChatCompletions, bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oai.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data llmv1.CompletionResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
