package server

import (
	"context"
	"log"
	"log/slog"

	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

func Serve(ctx context.Context) error {
	connectMux, err := connectrpc.NewConnectMux(connectrpc.Deps{})
	if err != nil {
		log.Fatal("connect rpc mux not created", err)
		return err
	}

	connectsrv := httpserver.New(connectMux)

	// wait for termination signal
	wait := utils.GracefulShutdown(ctx, connectsrv.Notify(), utils.DefaultShutdownTimeout, map[string]utils.Operation{
		"server": func(newCtx context.Context) error {
			return connectsrv.Shutdown()
		},
	})
	<-wait

	slog.Info("graceful shutdown complete")
	return nil
}
