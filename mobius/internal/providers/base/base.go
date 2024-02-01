package base

import (
	"context"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

type ProviderConfig struct {
	BaseURL         string
	ChatCompletions string
}

type ProviderInterface interface {
	GetName() string
	Validate() error
}

type ChatCompletionInterface interface {
	ProviderInterface
	ChatCompletion(context.Context, *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error)
}
