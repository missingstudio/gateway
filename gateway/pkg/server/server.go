package server

import (
	"context"
	"log/slog"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

func Serve(ctx context.Context, logger *slog.Logger, cfg *config.Config) error {
	connectMux, err := connectrpc.NewConnectMux(connectrpc.Deps{})
	if err != nil {
		logger.Error("connect rpc mux not created", err)
		return err
	}

	connectsrv := httpserver.New(connectMux, httpserver.WithAddr(cfg.Host, cfg.Port))

	// wait for termination signal
	wait := utils.GracefulShutdown(ctx, logger, connectsrv.Notify(), utils.DefaultShutdownTimeout, map[string]utils.Operation{
		"server": func(newCtx context.Context) error {
			return connectsrv.Shutdown()
		},
	})

	logger.Info("api server starting", "http-port", cfg.Port, "grpc-port", cfg.Port)
	<-wait

	logger.Info("graceful shutdown complete")
	return nil
}
