package base

import (
	"context"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}

type LLMProvider interface {
	ChatCompilation(ctx context.Context, ra *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error)
}
