package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/missingstudio/studio/backend/config"
	v1 "github.com/missingstudio/studio/backend/internal/api/v1"
	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/pkg/utils"
)

func Serve(ctx context.Context, logger *slog.Logger, cfg *config.Config) error {
	ingester := ingester.GetIngester(ctx, cfg.Ingester, logger)
	connectMux, err := connectrpc.NewConnectMux(v1.NewDeps(ingester))
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

	logger.Info("API server starting", "http-port", cfg.Port, "grpc-port", cfg.Port)
	logger.Info(fmt.Sprintf("🌈 AI Gateway is now running on http://localhost:%d", cfg.Port))
	<-wait

	logger.Info("graceful shutdown complete")
	return nil
}
