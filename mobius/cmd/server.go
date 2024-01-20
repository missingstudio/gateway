package cmd

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/MakeNowJust/heredoc"
	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"

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
			connectMux, err := connectrpc.NewConnectMux(connectrpc.Deps{})
			if err != nil {
				log.Fatal("connect rpc mux not created", err)
				return err
			}

			connectsrv := httpserver.New(connectMux, httpserver.WithAddr("127.0.0.1", "8080"))

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
			defer signal.Stop(interrupt)

			select {
			case s := <-interrupt:
				slog.Info("received interrupt signal", "signal", s.String())
			case err := <-connectsrv.Notify():
				slog.Error("got error from connect server", "error", err.Error())
			}

			if err := connectsrv.Shutdown(); err != nil {
				slog.Error("go error on connect server shutdown", "error", err.Error())
			}
			return nil
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}
