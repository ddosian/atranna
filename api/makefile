.PHONY: docs build run test

docs:
	swag init

build: docs
	go build .

run: docs
	go run .

test:
	go test ./...