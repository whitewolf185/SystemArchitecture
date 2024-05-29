
bin-deps:
	@echo -- 🛠️ build local depencies --
	go install golang.org/x/tools/cmd/stringer@latest

run: bin-deps
	go run ./cmd/main.go

run-docker: bin-deps
	go run ./cmd/main.go
