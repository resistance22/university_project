#!/bin/bash

export $(grep -v "^#" test.env | xargs -d '\n')

COMPOSE_FILE="compose.test.yaml"

docker compose -f $COMPOSE_FILE up -d

echo "Waiting for database to be healthy..."
docker compose -f $COMPOSE_FILE exec db bash -c 'until pg_isready -U testuser; do sleep 1; done'

docker compose -f $COMPOSE_FILE exec -it db psql -U postgres -c "create database project_database"

make migrattestup

echo "Running tests..."
go test --cover -v ./... --coverprofile=coverage.out -covermode=count

test_result=$?

docker compose -f $COMPOSE_FILE down

exit $test_result