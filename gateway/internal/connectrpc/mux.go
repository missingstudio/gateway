package connectrpc

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"

	"github.com/missingstudio/ai/gateway/internal/api"
	v1 "github.com/missingstudio/ai/gateway/internal/api/v1"
	"github.com/missingstudio/ai/protos/pkg/llm/v1/llmv1connect"
	"github.com/missingstudio/ai/protos/pkg/prompt/v1/promptv1connect"
)

func NewConnectMux(d *api.Deps) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	compress1KB := connect.WithCompressMinBytes(1024)
	v1Handler, err := v1.Register(d)
	if err != nil {
		return nil, fmt.Errorf("failed to create handler: %w", err)
	}

	mux.Handle("/", v1Handler)
	mux.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong\n")
	}))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(
			llmv1connect.LLMServiceName,
			promptv1connect.PromptRegistryServiceName,
		),
		compress1KB,
	))

	reflector := grpcreflect.NewStaticReflector(
		llmv1connect.LLMServiceName,
		promptv1connect.PromptRegistryServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector, compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, compress1KB))
	return mux, nil
}
