package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/internal/routes"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/internal/services"
)

func initDB() *pgx.Conn {
	connString := "postgres://" + os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@" + os.Getenv("DATABASE_HOST") + ":5432/" + os.Getenv("DATABASE_NAME")
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Connected to Database...")
	return conn
}

func main() {
	db := initDB()
	srv := services.NewService(db)
	mux := chi.NewMux()
	routes.MountApp(mux, srv)
	log.Println("Server runnning on port 8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err.Error())
	}
}
