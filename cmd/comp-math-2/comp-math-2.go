package main

import (
	"comp-math-2/internal/config"
	"comp-math-2/internal/web"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.Get()

	if err != nil {
		panic(err)
	}

	server := web.New(cfg)

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	if err := server.Start(); err != nil {
		log.Info("server shutdown", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
	}
}
