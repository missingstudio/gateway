package azure

import (
	"context"
	"errors"

	"github.com/missingstudio/studio/backend/internal/providers/openai"
	"github.com/missingstudio/studio/backend/models"
)

func (az *azureProvider) ChatCompletion(ctx context.Context, payload *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) Models() []string {
	return openai.OpenAIModels
}
