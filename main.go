package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/resistance22/university_project/api"
	db "github.com/resistance22/university_project/db/sqlc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
