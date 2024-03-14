package azure

import (
	"context"
	"errors"

	"github.com/missingstudio/ai/gateway/core/chat"
	"github.com/missingstudio/ai/gateway/internal/provider/openai"
)

func (az *azureProvider) ChatCompletions(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) Models() []string {
	return openai.OpenAIModels
}
