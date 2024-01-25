package base

import (
	"context"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}

type ProviderInterface interface{}

type ChatCompilationInterface interface {
	ProviderInterface
	ChatCompilation(context.Context, *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error)
}
