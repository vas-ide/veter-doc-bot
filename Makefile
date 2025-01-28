.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o cmd/bot/main.go