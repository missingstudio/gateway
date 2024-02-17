package openai

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/models"
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

func (oai *openAIProvider) ChatCompletion(ctx context.Context, payload []byte) (*http.Response, error) {
	client := requester.NewHTTPClient()
	requestURL := fmt.Sprintf("%s%s", oai.config.BaseURL, oai.config.ChatCompletions)
	req, _ := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(payload))

	connectionConfigMap := oai.conn.GetHeaders([]string{
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

func (*openAIProvider) Models() []string {
	return OpenAIModels
}
