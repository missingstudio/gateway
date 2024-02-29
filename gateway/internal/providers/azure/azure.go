package azure

import (
	"context"
	"errors"

	"github.com/missingstudio/studio/backend/core/chat"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
)

func (az *azureProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) Models() []string {
	return openai.OpenAIModels
}
