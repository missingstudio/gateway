package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	DefaultShutdownTimeout = 20 * time.Second
)

// Operation is a cleanup function on shutting down
type Operation func(ctx context.Context) error

func GracefulShutdown(ctx context.Context, notifier <-chan error, timeout time.Duration, ops map[string]Operation) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		defer signal.Stop(interrupt)

		select {
		case s := <-interrupt:
			slog.Info("received interrupt signal", "signal", s.String())
		case err := <-notifier:
			slog.Error("got error from server", "error", err.Error())
		}

		slog.Info("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			panic(fmt.Sprintf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds()))
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup
		// Do the operations asynchronously to save time
		for innerKey, innerOp := range ops {
			wg.Add(1)
			func() {
				defer wg.Done()

				slog.Info(fmt.Sprintf("cleaning up: %s", innerKey))
				if err := innerOp(ctx); err != nil {
					panic(fmt.Sprintf("%s: clean up failed: %s", innerKey, err.Error()))
				}

				slog.Info(fmt.Sprintf("%s was shutdown gracefully", innerKey))
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}
