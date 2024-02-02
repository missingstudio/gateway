package deepinfra

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/pkg/requester"
)

func (deepinfra *DeepinfraProvider) ChatCompletion(ctx context.Context, payload []byte) (*http.Response, error) {
	client := requester.NewHTTPClient()
	requestURL := fmt.Sprintf("%s%s", deepinfra.Config.BaseURL, deepinfra.Config.ChatCompletions)
	req, _ := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", deepinfra.APIKey))

	return client.SendRequestRaw(req)
}
