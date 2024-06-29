package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/resistance22/micorsales/api"
	db "github.com/resistance22/micorsales/db/sqlc"
)

func main() {
	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(dbContext)
	store := db.NewPGXStore(conn)
	server := api.NewServer(store)
	if err := server.Start("0.0.0.0:4321"); err != nil {
		log.Fatal(err)
	}
}
