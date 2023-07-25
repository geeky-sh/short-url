package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	handlers "shorturl/handlers"
	"shorturl/repositories"
	"shorturl/usecases"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	/*
		1. Introduce chi router
		2. Add health check
		3. Add database and ping value to the health check
		4. env parser and use the same
		5. Start adding APIs.
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env")
	}

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))

	hh := handlers.NewHealthHandler(db)
	r.Mount("/", hh.Routes())

	uh := handlers.NewURLhandler(
		usecases.NewURLUsecase(
			repositories.NewURLRepository(db),
		),
	)
	r.Mount("/", uh.Routes())

	http.ListenAndServe(":3333", r)
}
