package azure

import (
	"context"
	"errors"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (az *AzureProvider) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	return nil, errors.New("Not yet implemented")
}
