package anyscale

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/pkg/requester"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (anyscale *AnyscaleProvider) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	payload, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}

	client := requester.NewHTTPClient()
	requestURL := fmt.Sprintf("%s%s", anyscale.Config.BaseURL, anyscale.Config.ChatCompletions)
	req, _ := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", anyscale.APIKey))

	resp, err := client.Do(req)
	fmt.Println(resp)
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
