package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ss-ayan/ayan/config"
	"github.com/ss-ayan/ayan/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load("ayan.yaml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	srv, err := server.New(cfg)
	if err != nil {
		log.Fatalf("building server: %v", err)
	}
	defer srv.Close()

	slog.Info("ayan server starting", "port", cfg.Server.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
