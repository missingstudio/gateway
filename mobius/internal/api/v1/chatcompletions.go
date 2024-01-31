package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/common/errors"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (s *V1Handler) ChatCompletions(
	ctx context.Context,
	req *connect.Request[llmv1.CompletionRequest],
) (*connect.Response[llmv1.CompletionResponse], error) {
	provider, err := providers.GetProvider(ctx, req.Header())
	if err != nil {
		return nil, errors.New(err)
	}

	completionProvider, ok := provider.(base.ChatCompilationInterface)
	if !ok {
		return nil, errors.NewInternalError("provider don't have chat compilation capabilities")
	}

	data, err := completionProvider.ChatCompilation(ctx, req.Msg)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(data), nil
}
