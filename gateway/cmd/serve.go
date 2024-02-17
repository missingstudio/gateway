package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/api"
	"github.com/missingstudio/studio/backend/internal/connections"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
	"github.com/missingstudio/studio/backend/internal/server"
	"github.com/missingstudio/studio/common/logger"
	"github.com/redis/go-redis/v9"
)

func Serve(cfg *config.Config) error {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancelFunc()

	logger := logger.New(cfg.Log.Json, logger.WithLevel(slog.Level(cfg.Log.Level)))
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

	providerService := providers.NewService()
	connectionService := connections.NewService()

	deps := api.NewDeps(logger, ingester, rl, providerService, connectionService)

	if err := server.Serve(ctx, logger, cfg.App, deps); err != nil {
		logger.Error("error starting server", "error", err)
		return err
	}

	return nil
}
