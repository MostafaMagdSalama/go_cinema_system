package main

import (
	"cinema_system/config/db"
	ENVconfig "cinema_system/config/environment"
	"cinema_system/internal/customer"
	"cinema_system/internal/movie"
	"cinema_system/internal/seat"
	"cinema_system/internal/show"
	"cinema_system/internal/theatre"
	"cinema_system/internal/ticket"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Print("application running")

	// load env vars
	cfg, err := ENVconfig.ReadEnvironmentVars("./.env")
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
		return
	}

	// setup DB
	fmt.Println("outside db setup")
	db, err := db.Setup_DB(cfg.DB_CONNECTION)
	if err != nil {
		log.Fatal("DB connection failed: ", err)
		return
	}
	defer db.Close()
	r := chi.NewRouter()
	movieRepo := movie.NewPostgresRepository(db)
	movieService := movie.NewService(movieRepo)
	movieHandler := movie.NewHandler(movieService)
	movieHandler.RegisterRoutes(r)

	showRepo := show.NewPostgresRepository(db)
	showService := show.NewService(showRepo)
	showHandler := show.NewHandler(showService)
	showHandler.RegisterRoutes(r)

	theatreRepo := theatre.NewPostgresRepository(db)
	theatreService := theatre.NewService(theatreRepo)
	theatreHandler := theatre.NewHandler(theatreService)
	theatreHandler.RegisterRoutes(r)

	seatRepo := seat.NewPostgresRepository(db)
	seatService := seat.NewService(seatRepo)
	seatHandler := seat.NewHandler(seatService)
	seatHandler.RegisterRoutes(r)

	customerRepo := customer.NewPostgresRepository(db)
	customerService := customer.NewService(customerRepo)
	customerHandler := customer.NewHandler(customerService)
	customerHandler.RegisterRoutes(r)

	ticketRepo := ticket.NewPostgresRepository(db)
	ticketService := ticket.NewService(ticketRepo)
	ticketHandler := ticket.NewHandler(ticketService)
	ticketHandler.RegisterRoutes(r)


	if err != nil {
		log.Fatal(err)
	}

	go shutdown(context.Background())
	http.ListenAndServe(":"+cfg.PORT, r)

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
