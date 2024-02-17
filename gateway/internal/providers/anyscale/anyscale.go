package anyscale

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/models"
	"github.com/missingstudio/studio/backend/pkg/requester"
)

func (anyscale *anyscaleProvider) ChatCompletion(ctx context.Context, payload []byte) (*http.Response, error) {
	client := requester.NewHTTPClient()
	requestURL := fmt.Sprintf("%s%s", anyscale.config.BaseURL, anyscale.config.ChatCompletions)
	req, _ := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(payload))

	connectionConfigMap := anyscale.conn.GetHeaders([]string{
		models.AuthorizationHeader,
	})

	var authorizationHeader string
	if val, ok := connectionConfigMap[models.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)

	return client.SendRequestRaw(req)
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
