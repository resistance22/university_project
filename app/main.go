package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	config "github.com/resistance22/university_project/Config"
	"github.com/resistance22/university_project/api"
	db "github.com/resistance22/university_project/db/sqlc"
)

func main() {
	envConfig, err := config.LoadConfig("..", "dev")

	if err != nil {
		log.Fatal(err.Error())
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), config.DBConfig(envConfig.DBUrl))

	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	store := db.NewPGXStore(connPool)
	server := api.NewServer(&envConfig, store)
	if err := server.Start(fmt.Sprintf("%s:%d", envConfig.ServerAddress, envConfig.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
