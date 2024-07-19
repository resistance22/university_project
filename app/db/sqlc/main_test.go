package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	config "github.com/resistance22/university_project/Config"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../..", "test")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, config.DBUrl)
	if err != nil {
		log.Fatal("Connection To DB Failed!")
	}
	testQueries = New(conn)
	code := m.Run()
	os.Exit(code)
}
