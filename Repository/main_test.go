package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	config "github.com/resistance22/university_project/Config"
	db "github.com/resistance22/university_project/db/sqlc"
)

var UserRepo *UserRepository

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("..", "test")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, config.DBUrl)
	if err != nil {
		log.Fatal("Connection To DB Failed!")
	}
	store := db.NewPGXStore(conn)
	UserRepo = &UserRepository{
		store: store,
	}
	code := m.Run()
	os.Exit(code)
}
