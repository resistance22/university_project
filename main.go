package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	server "github.com/resistance22/micorsales/api"
)

func main() {
	dbContext := context.Background()
	conn, err := pgxpool.New(dbContext, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	apiServer := server.NewServer()
	apiServer.Start("0.0.0.0:3000")
}
