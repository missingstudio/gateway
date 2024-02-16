package anyscale

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/pkg/requester"
)

func (anyscale *anyscaleProvider) ChatCompletion(ctx context.Context, payload []byte) (*http.Response, error) {
	client := requester.NewHTTPClient()
	requestURL := fmt.Sprintf("%s%s", anyscale.Config.BaseURL, anyscale.Config.ChatCompletions)
	req, _ := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", anyscale.APIKey))

	return client.SendRequestRaw(req)
}

func (anyscale *anyscaleProvider) GetModels() []interface{} {
	return []interface{}{
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
