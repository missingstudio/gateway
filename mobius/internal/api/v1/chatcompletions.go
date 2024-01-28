package v1

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (s *V1Handler) ChatCompletions(
	ctx context.Context,
	req *connect.Request[llmv1.CompletionRequest],
) (*connect.Response[llmv1.CompletionResponse], error) {
	provider, err := providers.GetProvider(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	completionProvider, ok := provider.(base.ChatCompilationInterface)
	if !ok {
		return nil, connect.NewError(connect.CodeUnimplemented, errors.New("method not implemented"))
	}

	data, err := completionProvider.ChatCompilation(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(data), nil
}
