package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"connectrpc.com/connect"

	v1 "github.com/missingstudio/studio/backend/internal/api/v1"
	"github.com/missingstudio/studio/backend/internal/errors"
	"github.com/missingstudio/studio/backend/internal/interceptor"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
	"github.com/zeebo/assert"

	"github.com/stretchr/testify/require"
)

func TestGatewayServer(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(llmv1connect.NewLLMServiceHandler(
		&v1.V1Handler{},
		connect.WithInterceptors(interceptor.HeadersInterceptor()),
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
			assert.True(t, strings.Contains(err.Error(), errors.ErrRequiredHeaderNotExit.Error()))
		}
	})
}
