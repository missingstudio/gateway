package azure

import (
	"context"
	"errors"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/internal/providers/openai"
)

func (az *azureProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) Models() []string {
	return openai.OpenAIModels
}
