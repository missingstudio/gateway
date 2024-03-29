package cmd

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/missingstudio/ai/gateway/config"
	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/api"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	iprovider "github.com/missingstudio/ai/gateway/internal/provider"
	"github.com/missingstudio/ai/gateway/internal/ratelimiter"
	"github.com/missingstudio/ai/gateway/internal/storage/postgres"
	"github.com/missingstudio/ai/gateway/pkg/database"
	"github.com/missingstudio/common/logger"
	"github.com/missingstudio/common/rest"
	"github.com/redis/go-redis/v9"
)

func Serve(cfg *config.Config) *api.API {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancelFunc()

	restConfig := &rest.Config{}
	restConfig.SetDefaults()

	logger := logger.New(cfg.Log.Json, logger.WithLevel(slog.Level(cfg.Log.Level)))
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("failed to init redis connection")
	}

	rate := ratelimiter.NewRate(cfg.Ratelimiter.DurationInSecond, cfg.Ratelimiter.NumberOfRequests)
	ingester := ingester.GetIngesterWithDefault(ctx, cfg.Ingester, logger)
	ratelimiter := ratelimiter.NewRateLimiter(rdb, logger, rate, cfg.Ratelimiter.Type)

	// prefer use pgx instead of lib/pq for postgres to catch pg error
	if cfg.Postgres.Driver == "postgres" {
		cfg.Postgres.Driver = "pgx"
	}

	dbClient, err := database.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to setup db: %v", err)
	}

	apikeyRepository := postgres.NewAPIKeyRepository(dbClient)
	apikeyService := apikey.NewService(apikeyRepository)

	promptRepository := postgres.NewPromptRepository(dbClient)
	promptService := prompt.NewService(promptRepository)

	providerRepository := postgres.NewProviderRepository(dbClient)
	providerService := provider.NewService(providerRepository)

	iProviderService := iprovider.NewService()

	return &api.API{
		Logger:           logger,
		DBClient:         dbClient,
		Ingester:         ingester,
		RestConfig:       restConfig,
		RateLimiter:      ratelimiter,
		APIKeyService:    apikeyService,
		PromptService:    promptService,
		ProviderService:  providerService,
		IProviderService: iProviderService,
		RequestTimeout:   cfg.App.RequestTimeout,
	}
}
