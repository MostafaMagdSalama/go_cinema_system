package main

import (
	config "cinema_system/config/db"
	ENVconfig "cinema_system/config/environment"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

func main() {

	log.Print("application running")

	// load env vars
	cfg, err := ENVconfig.ReadEnvironmentVars(".env")
	if err != nil {
		errors.Wrap(err, "loadig env vars")
	}

	// setup DB
	db, err := config.Setup_DB(cfg.DB_CONNECTION)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	go shutdown(context.Background())

}

func shutdown(ctx context.Context) {
	// graceful shutdown
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	<-ctx.Done()

	log.Print("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform any cleanup tasks here
	log.Print("performing cleanup tasks")

	// Wait for cleanup to complete or timeout
	<-shutdownCtx.Done()
	log.Print("cleanup completed or timeout reached")
}
