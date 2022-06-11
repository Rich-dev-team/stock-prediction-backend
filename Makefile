migrateUp:
	migrate -path ./db/migration/ -database "postgresql://postgres:rootadmin@34.80.203.194:5432/stockdb" -verbose up
migrateDown:
	migrate -path ./db/migration/ -database "postgresql://postgres:rootadmin@34.80.203.194:5432/stockdb" -verbose down

.PHONY: migrateUp migrateDown