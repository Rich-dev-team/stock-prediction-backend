
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

cleanDirty:
	migrate -path ./db/migration/ -database "postgresql://root:secret@localhost:5432/stockdb?sslmode=disable" -verbose force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

connDB:
	docker exec -it postgres12 psql -U root -d stockdb

server:
	go run main.go

PHONY= postgres createdb dropdb migrateUp migrateDown sqlc test connDB server

.PHONY: $(PHONY)