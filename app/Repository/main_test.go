package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	config "github.com/resistance22/university_project/Config"
	db "github.com/resistance22/university_project/db/sqlc"
)

var UserRepo *UserRepository

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..", "test")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbContext := context.Background()
	fmt.Println(config.DBUrl)
	conn, err := pgx.Connect(dbContext, config.DBUrl)
	if err != nil {

		log.Fatal(err.Error())
	}
	store := db.NewPGXStore(conn)
	UserRepo = &UserRepository{
		store: store,
	}
	code := m.Run()
	os.Exit(code)
}
