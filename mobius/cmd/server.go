package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/MakeNowJust/heredoc"
	"github.com/missingstudio/studio/backend/pkg/server"

	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Server management",
		Long:    "Server management commands.",
		Example: heredoc.Doc(`
			$ mobius server start
		`),
	}

	cmd.AddCommand(serverStartCommand())
	return cmd
}

func serverStartCommand() *cobra.Command {
	var configFile string

	c := &cobra.Command{
		Use:     "start",
		Short:   "Start server and proxy default on port 8080",
		Example: "frontier server start",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
			defer cancelFunc()

			return server.Serve(ctx)
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}
