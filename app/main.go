package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	config "github.com/resistance22/university_project/Config"
	"github.com/resistance22/university_project/api"
	db "github.com/resistance22/university_project/db/sqlc"
)

func main() {
	config, err := config.LoadConfig("..", "dev")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, config.DBUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(dbContext)
	store := db.NewPGXStore(conn)
	server := api.NewServer(&config, store)
	if err := server.Start(fmt.Sprintf("%s:%d", config.ServerAddress, config.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
