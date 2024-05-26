package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	dbContext := context.Background()
	conn, err := pgx.Connect(dbContext, dbSource)
	if err != nil {
		log.Fatal("Connection To DB Failed!")
	}
	testQueries = New(conn)
	code := m.Run()
	os.Exit(code)
}
