package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	v1 "github.com/missingstudio/ai/gateway/internal/api/v1"
	llmv1 "github.com/missingstudio/protos/pkg/llm/v1"
	"github.com/missingstudio/protos/pkg/llm/v1/llmv1connect"

	"github.com/stretchr/testify/require"
)

func TestGatewayServer(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(llmv1connect.NewLLMServiceHandler(
		&v1.V1Handler{},
		connect.WithInterceptors(),
	))

	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := llmv1connect.NewLLMServiceClient(
		server.Client(),
		server.URL,
	)

	grpcClient := llmv1connect.NewLLMServiceClient(
		server.Client(),
		server.URL,
		connect.WithGRPC(),
	)

	clients := []llmv1connect.LLMServiceClient{
		connectClient,
		grpcClient,
	}

	t.Run("chat completions: shoud provide provider in headers", func(t *testing.T) {
		for _, client := range clients {

			req := connect.NewRequest(&llmv1.ChatCompletionRequest{})
			_, err := client.ChatCompletions(context.Background(), req)
			require.NotNil(t, err)
		}
	})
}
