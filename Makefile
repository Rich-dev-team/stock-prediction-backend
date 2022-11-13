
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

migrateup1:
	migrate -path ./db/migration/ -database "postgresql://root:secret@localhost:5432/stockdb?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path ./db/migration/ -database "postgresql://root:secret@localhost:5432/stockdb?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

connDB:
	docker exec -it postgres12 psql -U root -d stockdb

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package=mockdb -destination db/mock/store.go github.com/Rich-dev-team/stock-prediction-backend/db/sqlc Store

PHONY= postgres createdb dropdb migrateUp migrateDown migrateUp1 migrateDown1 sqlc test connDB server mock

.PHONY: $(PHONY)