
postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root stockdb

dropdb:
	docker exec -it postgres12 dropdb stockdb

migrateup:
	migrate -path ./db/migration/ -database "postgresql://root:secret@localhost:5432/stockdb?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration/ -database "postgresql://root:secret@localhost:5432/stockdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

PHONY= postgres createdb dropdb migrateUp migrateDown sqlc

.PHONY: $(PHONY)