package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/solace06/cron-runner/api"
	"github.com/solace06/cron-runner/job/config"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//setup database

	//setup router
	router := api.NewRouter()

	//setup server
	slog.Info("started server", slog.String("address", cfg.Address))
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shut down the server", slog.String("error: ", err.Error()))
	}
	slog.Info("server shutdown successfully")
}