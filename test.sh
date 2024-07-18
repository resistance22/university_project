#!/bin/bash

export $(grep -v "^#" .env.test | xargs -d '\n')

docker compose -f compose.test.yaml up -d

echo "Waiting for database to be healthy..."
docker compose exec db bash -c 'until pg_isready -U testuser; do sleep 1; done'

docker compose exec -it db psql -U postgres -c "create database project_database"

make migrattestup

echo "Running tests..."
go test --cover -v ./...

test_result=$?

docker compose down

exit $test_result