MIGRATION_DIR="./migrations"
LOCAL_DB=domain_service
GOOSE_DRIVER=mysql
DBSTRING="user=new_user dbname=$(LOCAL_DB) sslmode=disable"
LOCAL_DB_DSN="mysql://new_user:password@localhost:3306/$(LOCAL_DB)?sslmode=disable"

DOCKER_DB_DSN:=postgres://postgres:postgres@db:5432/$(LOCAL_DB)

bin-deps:
	@echo -- 🛠️ build local depencies --
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/go-jet/jet/v2/cmd/jet@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest

docker-bin-deps:
	@echo -- 🛠️ build docker local depencies --
	go install github.com/pressly/goose/v3/cmd/goose@latest

db-reset:
	@echo -- 🛢 reset database $(LOCAL_DB) --
	psql -U postgres -c "DROP DATABASE IF EXISTS $(LOCAL_DB) WITH (FORCE);"
	psql -U postgres -c "CREATE DATABASE $(LOCAL_DB);"
	goose -dir $(MIGRATION_DIR) $(GOOSE_DRIVER) $(DBSTRING) up
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
	goose -dir $(MIGRATION_DIR) $(GOOSE_DRIVER) $(DBSTRING) up
docker-migrate-up: docker-bin-deps
	@echo -- 🛢 docker migrate up --
	goose -dir $(MIGRATION_DIR) -allow-missing postgres $(DOCKER_DB_DSN) up

run: docker-migrate-up
	go run ./cmd/main.go
