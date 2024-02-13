package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/missingstudio/studio/backend/config"
	v1 "github.com/missingstudio/studio/backend/internal/api/v1"
	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func Serve(ctx context.Context, logger *slog.Logger, cfg *config.Config) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Warn("failed to init redis connection")
		os.Exit(1)
	}

	rl := ratelimiter.NewRateLimiter(cfg.Ratelimiter, logger, cfg.Ratelimiter.Type, rdb)
	ingester := ingester.GetIngesterWithDefault(ctx, cfg.Ingester, logger)

	connectMux, err := connectrpc.NewConnectMux(v1.NewDeps(ingester, rl))
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
