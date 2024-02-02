package connectrpc

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	v1 "github.com/missingstudio/studio/backend/internal/api/v1"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
)

type Deps struct{}

func NewConnectMux(d Deps) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	compress1KB := connect.WithCompressMinBytes(1024)
	v1Handler, err := v1.Register()
	if err != nil {
		return nil, fmt.Errorf("failed to create handler: %w", err)
	}

	mux.Handle("/", v1Handler)
	mux.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong\n")
	}))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(llmv1connect.LLMServiceName),
		compress1KB,
	))

	reflector := grpcreflect.NewStaticReflector(llmv1connect.LLMServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector, compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, compress1KB))

	return mux, nil
}
