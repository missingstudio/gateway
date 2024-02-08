package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/server"
	"github.com/missingstudio/studio/common/logger"
)

func Serve(cfg *config.Config) error {
	logger := logger.New(cfg.LogFormatJson, nil)

	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancelFunc()

	if err := server.Serve(ctx, logger, cfg); err != nil {
		logger.Error("error starting server", "error", err)
		return err
	}

	return nil
}
