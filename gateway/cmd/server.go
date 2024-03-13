package cmd

import (
	"fmt"
	"log/slog"

	"github.com/MakeNowJust/heredoc"
	"github.com/missingstudio/ai/gateway/config"
	"github.com/missingstudio/common/logger"

	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Server management",
		Long:    "Server management commands.",
		Example: heredoc.Doc(`
			$ gateway server start
		`),
	}

	cmd.AddCommand(serverInitCommand())
	cmd.AddCommand(serverStartCommand())
	cmd.AddCommand(serverMigrateCommand())
	cmd.AddCommand(serverMigrateRollbackCommand())
	return cmd
}

func serverInitCommand() *cobra.Command {
	var configFile string
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialize server",
		Long: heredoc.Doc(`
			Initializing server. Creating a sample of missing studio server config.
			Default: ./config.yaml
		`),
		Example: "missing studio server init",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Init(configFile); err != nil {
				return err
			}

			fmt.Printf("server config created: %v\n", configFile)
			return nil
		},
	}

	c.Flags().StringVarP(&configFile, "output", "o", "./config.yaml", "Output config file path")
	return c
}

func serverStartCommand() *cobra.Command {
	var configFile string

	c := &cobra.Command{
		Use:     "start",
		Short:   "Start server and proxy default on port 8080",
		Example: "missing studio server start",
		RunE: func(cmd *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}
			return Serve(appConfig)
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}

func serverMigrateCommand() *cobra.Command {
	var configFile string

	c := &cobra.Command{
		Use:     "migrate",
		Short:   "Run DB Schema Migrations",
		Example: "missing studio server migrate",
		RunE: func(c *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}

			logger := logger.New(appConfig.Log.Json, logger.WithLevel(slog.Level(appConfig.Log.Level)))
			logger.Info("missing studio is migrating", "version", config.Version)

			if err = RunMigrations(logger, appConfig.Postgres); err != nil {
				logger.Error("error running migrations", "error", err)
				return err
			}

			logger.Info("missing studio  migration complete")
			return nil
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}

func serverMigrateRollbackCommand() *cobra.Command {
	var configFile string

	c := &cobra.Command{
		Use:     "migrate-rollback",
		Short:   "Run DB Schema Migrations Rollback to last state",
		Example: "missing studio migrate-rollback",
		RunE: func(c *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}
			logger := logger.New(appConfig.Log.Json, logger.WithLevel(slog.Level(appConfig.Log.Level)))
			logger.Info("missing studio is migrating", "version", config.Version)

			if err = RunRollback(logger, appConfig.Postgres); err != nil {
				logger.Error("error running migrations rollback", "error", err)
				return err
			}

			logger.Info("missing studio migration rollback complete")
			return nil
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}
