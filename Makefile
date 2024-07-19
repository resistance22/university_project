postgres:
	docker compose -f ./compose.yaml up -d
createdb:
	docker exec -it project_db createdb --username=postgres --owner=postgres project_database
dropdb:
	docker exec -it project_db dropdb --username=postgres project_database
createmigration:
	migrate create -ext sql -dir app/db/migration -seq new_tables
migrateup:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose up
migratedown:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose down
migrateupone:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose up 1
migratedownone:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/project_database?sslmode=disable" -verbose down 1
migrattestoneup:
	migrate -path app/db/migration -database ${DB_URL} -verbose up 1
migrattestonedown:
	migrate -path app/db/migration -database ${DB_URL} -verbose down 1
migrattestup:
	migrate -path app/db/migration -database ${DB_URL} -verbose up
migrattestdown:
	migrate -path app/db/migration -database ${DB_URL} -verbose down
build_image:
	docker build -t project_university .
sqlc:
	sqlc generate
test:
	go test --cover -v ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown test migrattestoneup migrattestonedown migrattestup migrattestdown build_image