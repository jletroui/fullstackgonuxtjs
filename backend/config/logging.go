package config

import (
	"log/slog"
	"os"
)

func ConfigureLogging() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelDebug)
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))
}
