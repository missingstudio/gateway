package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/core/connection"
	"github.com/missingstudio/studio/backend/core/prompt"
	"github.com/missingstudio/studio/backend/internal/api"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/backend/internal/ratelimiter"
	"github.com/missingstudio/studio/backend/internal/server"
	"github.com/missingstudio/studio/backend/internal/storage/postgres"
	"github.com/missingstudio/studio/backend/pkg/database"
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

	rate := ratelimiter.NewRate(cfg.Ratelimiter.DurationInSecond, cfg.Ratelimiter.NumberOfRequests)
	rl := ratelimiter.NewRateLimiter(rdb, logger, rate, cfg.Ratelimiter.Type)
	ingester := ingester.GetIngesterWithDefault(ctx, cfg.Ingester, logger)

	// prefer use pgx instead of lib/pq for postgres to catch pg error
	if cfg.Postgres.Driver == "postgres" {
		cfg.Postgres.Driver = "pgx"
	}

	dbc, err := database.New(cfg.Postgres)
	if err != nil {
		err = fmt.Errorf("failed to setup db: %w", err)
		return err
	}

	defer func() {
		logger.Debug("cleaning up db")
		if err := dbc.Close(); err != nil {
			logger.Warn("db cleanup failed", "err", err)
		}
	}()

	connectionRepository := postgres.NewConnectionRepository(dbc)
	connectionService := connection.NewService(connectionRepository)

	promptRepository := postgres.NewPromptRepository(dbc)
	promptService := prompt.NewService(promptRepository)

	providerService := providers.NewService()
	deps := api.NewDeps(logger, ingester, rl, providerService, connectionService, promptService)

	if err := server.Serve(ctx, logger, cfg.App, deps); err != nil {
		logger.Error("error starting server", "error", err)
		return err
	}

	return nil
}
