package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	greetingv1 "github.com/missingstudio/protos/pkg/greeting/v1"
	"github.com/missingstudio/protos/pkg/greeting/v1/greetingv1connect"
	"github.com/missingstudio/studio/backend/internal/connectrpc"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestMobiusSercer(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(greetingv1connect.NewGreetServiceHandler(
		&connectrpc.GreetServer{},
	))

	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := greetingv1connect.NewGreetServiceClient(
		server.Client(),
		server.URL,
	)

	grpcClient := greetingv1connect.NewGreetServiceClient(
		server.Client(),
		server.URL,
		connect.WithGRPC(),
	)

	clients := []greetingv1connect.GreetServiceClient{
		connectClient,
		grpcClient,
	}

	t.Run("greet", func(t *testing.T) {
		for _, client := range clients {
			result, err := client.Greet(context.Background(), connect.NewRequest(&greetingv1.GreetRequest{
				Name: "Dev",
			}))

			require.Nil(t, err)
			assert.True(t, len(result.Msg.Greeting) > 0)
		}
	})
}
