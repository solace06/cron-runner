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

	"github.com/solace06/cron-runner/job"
)

func main() {

	s, er := job.NewScope()
	if er != nil {
		log.Fatalf("failed to initialize scope: %v", er)
	}
	
	defer func() {
		if err := s.Close(); err != nil {
			slog.Error("error closing the database", slog.String("error: ", err.Error()))
		}
	}()

	s.Migrate()

	//setup router
	router := job.NewRouter(s)

	//setup server
	slog.Info("started server", slog.String("address", s.Cfg.Address))
	server := http.Server{
		Addr:    s.Cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server %v", err)
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
