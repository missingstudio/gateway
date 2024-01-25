package connectrpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
)

type Deps struct{}

func NewConnectMux(d Deps) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("validate interceptor not created: %w", err)
	}

	compress1KB := connect.WithCompressMinBytes(1024)
	services := []*vanguard.Service{
		vanguard.NewService(llmv1connect.NewLLMServiceHandler(
			&LLMServer{},
			compress1KB,
			connect.WithInterceptors(validateInterceptor),
		)),
	}
	transcoderOptions := []vanguard.TranscoderOption{
		vanguard.WithUnknownHandler(Custom404handler()),
	}

	transcoder, err := vanguard.NewTranscoder(services, transcoderOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transcoder: %w", err)
	}
	mux.Handle("/", transcoder)
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(llmv1connect.LLMServiceName),
		compress1KB,
	))

	reflector := grpcreflect.NewStaticReflector(llmv1connect.LLMServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector, compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, compress1KB))

	return mux, nil
}

type LLMServer struct {
	llmv1connect.UnimplementedLLMServiceHandler
}

func (s *LLMServer) ChatCompletions(
	ctx context.Context,
	req *connect.Request[llmv1.CompletionRequest],
) (*connect.Response[llmv1.CompletionResponse], error) {
	log.Println("Request headers: ", req.Header())

	provider, err := providers.GetProvider(req.Header())
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
