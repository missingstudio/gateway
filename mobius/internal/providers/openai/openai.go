package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

var OpenAIAPIURL = "https://api.openai.com/v1/chat/completions"

type OpenAI struct {
	APIKey string
}

func (oai OpenAI) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	payload, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "POST", OpenAIAPIURL, bytes.NewReader(payload))
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
