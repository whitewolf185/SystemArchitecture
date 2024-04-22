
bin-deps:
	@echo -- 🛠️ build local depencies --
	go install go.uber.org/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install golang.org/x/tools/cmd/stringer@latest

swag-gen: bin-deps
	@echo -- 😎 swagger generation --
	swag init -p snakecase --parseInternal --pd -g ./internal/app/service.go

run: swag-gen
	go run ./cmd/main.go
