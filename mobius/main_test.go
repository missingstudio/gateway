package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	"github.com/missingstudio/studio/backend/internal/connectrpc"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestMobiusSercer(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(llmv1connect.NewLLMServiceHandler(
		&connectrpc.LLMServer{},
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

	t.Run("chat completions", func(t *testing.T) {
		for _, client := range clients {
			result, err := client.ChatCompletions(context.Background(), connect.NewRequest(&llmv1.CompletionRequest{}))

			require.Nil(t, err)
			assert.True(t, len(result.Msg.Object) > 0)
		}
	})
}
