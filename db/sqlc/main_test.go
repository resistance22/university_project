package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(os.Getenv("DB_URL"))

	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Connection To DB Failed!")
	}
	testQueries = New(conn)
	code := m.Run()
	os.Exit(code)
}
