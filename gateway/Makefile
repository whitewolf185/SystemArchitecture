
bin-deps:
	@echo -- 🛠️ build local depencies --
	go install go.uber.org/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install golang.org/x/tools/cmd/stringer@latest

run: bin-deps
	go run ./cmd/main.go
