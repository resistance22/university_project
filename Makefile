postgres:
	docker compose up -d
createdb:
	docker exec -it project_db createdb --username=postgres --owner=postgres project_database
dropdb:
	docker exec -it project_db dropdb --username=postgres project_database
createmigration:
	migrate create -ext sql -dir db/migration -seq new_tables
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose down
migrateupone:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose up 1
migratedownone:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose down 1
migrattestoneup:
	migrate -path db/migration -database ${DB_URL} -verbose up 1
migrattestonedown:
	migrate -path db/migration -database ${DB_URL} -verbose down 1
migrattestup:
	migrate -path db/migration -database ${DB_URL} -verbose up
migrattestdown:
	migrate -path db/migration -database ${DB_URL} -verbose down

sqlc:
	sqlc generate
test:
	go test --cover -v ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown test migrattestoneup migrattestonedown migrattestup migrattestdown