package mslogger

import (
	"log/slog"
	"os"
)

func New(formatAsJson bool, opts *slog.HandlerOptions) *slog.Logger {
	if formatAsJson {
		return slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}
