package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	fs embed.FS
)

func main() {
	databaseUrl := "postgres://aash:@localhost:5432/shorturl"
	db, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Could not connect to DB. %v\n", err)
	}

	d, err := iofs.New(fs, "db/migrations")
	if err != nil {
		log.Fatalf("Could not open file %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
	if err != nil {
		log.Fatalf("Could not create migration instance. %v\n", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("Error while running migrations %v", err)
	}

}
