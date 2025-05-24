package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("run() returned error: ", err.Error())
	}
}

func run() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	defer stop()

	slog.Info("Loading config from env")
	cfg, err := loadCfgFromEnv()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	slog.Info("Config loaded")

	select {
	case <-ctx.Done():
		return nil
	default:
	}

	http.ListenAndServe(":"+cfg.Port, nil)
	return nil
}
