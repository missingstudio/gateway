package providers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

type LLMProvider interface {
	ChatCompilation(ctx context.Context, ra *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error)
}

func NewLLMProvider(headers http.Header) (LLMProvider, error) {
	provider := headers.Get("x-ms-provider")
	if provider == "" {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("provider not found"))
	}
	authHeader := headers.Get("Authorization")
	accessToken := strings.Replace(authHeader, "Bearer ", "", 1)

	return &openai.OpenAI{APIKey: accessToken}, nil
}
