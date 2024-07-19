#!/bin/bash

export $(grep -v "^#" test.env | xargs -d '\n')

COMPOSE_FILE="compose.test.yaml"

echo "DB_URL:"
echo $DB_URL

docker compose -f $COMPOSE_FILE up -d

echo "Waiting for database to be healthy..."
docker compose -f $COMPOSE_FILE exec test_db bash -c 'until pg_isready -U test_user; do sleep 1; done'
# 
docker compose -f $COMPOSE_FILE exec test_db psql -U test_user -c "create database if not exists test_db"

make migrattestup

echo "Running tests..."
go test --cover -v ./... --coverprofile=coverage.out -covermode=count

test_result=$?

docker compose -f $COMPOSE_FILE down

exit $test_result