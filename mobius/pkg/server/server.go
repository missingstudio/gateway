package server

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/missingstudio/studio/backend/internal/connectrpc"
	"github.com/missingstudio/studio/backend/internal/httpserver"
)

func Serve(ctx context.Context) error {
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
}
