package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"url-short/internal/config"
	pgrepo "url-short/internal/repo/pgRepo"
	"url-short/internal/service"
	httpserver "url-short/internal/transport/http_server"
	"url-short/pkg/pg"

	"github.com/gorilla/mux"
)

const (
	local = "local"
	dev = "develop"
)

func main() {
	if err := run(); err !=nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.Deal()
	if err != nil {
		return err
	}

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("logger enabled")

	pg, err := pg.Deal(cfg.Dsn)
	if err != nil {
		return err
	}

	log.Info("database connect succesfull")

	urlRepo := pgrepo.New(pg)
	urlService := service.New(log, urlRepo)

	httpServer := httpserver.New(urlService)

	router := mux.NewRouter()

	router.HandleFunc("/create", httpServer.Create).Methods(http.MethodPost)
	router.HandleFunc("/{alias}", httpServer.RedirectByAlias).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: router,
		Addr: cfg.Server.Addr,
	}

	log.Info("server starting...")

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed server: %w", err)
	}


	signChan := make(chan os.Signal, 1)

	go func(){
		signal.Notify(signChan, syscall.SIGINT, os.Interrupt)
		<-signChan
		
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		srv.Shutdown(ctx)
		pg.Close()
		os.Exit(1)
	}()
	
	return nil
}


func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch(env) {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

