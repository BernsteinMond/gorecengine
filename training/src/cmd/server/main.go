package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
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
	_, err = loadCfgFromEnv()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	slog.Info("Config loaded")

	select {
	case <-ctx.Done():
		return nil
	default:
	}

	stopWg := sync.WaitGroup{}

	stopWg.Add(1)
	go func(ctx context.Context) {
		defer stopWg.Done()
		httpSrvErr := launchHTTPServer(ctx, nil) // TODO: replace with *http.Server
		if httpSrvErr != nil {
			slog.Error("launch http server returned error", slog.String("error", httpSrvErr.Error()))
		}
	}(ctx)

	return nil
}

func launchHTTPServer(ctx context.Context, server *http.Server) (err error) {
	var httpServerShutDownError error
	defer func() {
		err = errors.Join(err, httpServerShutDownError)
	}()

	shutDownDone := make(chan struct{})
	go func(ctx context.Context) {
		<-ctx.Done()

		slog.Info("Shutting down http server")
		httpServerShutDownError = server.Shutdown(ctx)
		slog.Info("Http server shut down")

		close(shutDownDone)
	}(ctx)

	select {
	case <-ctx.Done():
		return nil
	default:
	}

	slog.Info("Starting http server", slog.String("addr", server.Addr))
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen on %s: %w", server.Addr, err)
	}

	<-shutDownDone
	return nil
}
