package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	handlers "shorturl/handlers"
	"shorturl/repositories"
	"shorturl/usecases"
	"shorturl/utils/session"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env")
	}

	db, err := sql.Open("sqlite3", "main.db?mode=rwc")
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	sess := session.Init()

	v := validator.New()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))
	r.Use(cors.Handler(cors.Options{AllowedHeaders: []string{"*"}}))

	hh := handlers.NewHealthHandler(db)
	r.Mount("/metrics", hh.Routes())

	ush := handlers.NewUserHandler(
		usecases.NewUserUsecase(
			repositories.NewUserRepository(db),
		), v, sess,
	)
	r.Mount("/users", ush.Routes())

	uh := handlers.NewURLhandler(
		usecases.NewURLUsecase(
			repositories.NewURLRepository(db),
		), v, sess,
	)
	r.Mount("/", uh.Routes())

	http.ListenAndServe(":4000", r)
}
