package azure

import (
	"context"
	"errors"

	"github.com/missingstudio/studio/backend/internal/providers/openai"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (az *azureProvider) ChatCompletion(ctx context.Context, cr *llmv1.ChatCompletionRequest) (*llmv1.ChatCompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) Models() []string {
	return openai.OpenAIModels
}
