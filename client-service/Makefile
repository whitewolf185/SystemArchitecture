MIGRATION_DIR="./migrations"
LOCAL_DB=client_service
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="user=postgres dbname=$(LOCAL_DB) sslmode=disable"
LOCAL_DB_DSN="postgres://postgres:qwerty@localhost:5432/$(LOCAL_DB)?sslmode=disable"

bin-deps:
	@echo -- 🛠️ build local depencies --
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/go-jet/jet/v2/cmd/jet@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest

db-reset:
	@echo -- 🛢 reset database $(LOCAL_DB) --
	psql -U postgres -c "DROP DATABASE IF EXISTS $(LOCAL_DB) WITH (FORCE);"
	psql -U postgres -c "CREATE DATABASE $(LOCAL_DB);"
	goose -dir $(MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
# generate schema
	pg_dump -s --username=postgres --dbname=$(LOCAL_DB) --file=schema.sql

jet: bin-deps
	@echo -- 🛢 jet generation --
	jet -dsn $(LOCAL_DB_DSN) -path=./internal/gen

swag-gen: bin-deps
	@echo -- 😎 swagger generation --
	swag init -p snakecase --parseInternal --pd -g ./internal/app/service.go

migrate-up: bin-deps
	@echo -- 🛢 migrate up --
	goose -dir $(MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
