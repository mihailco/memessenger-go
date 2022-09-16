.PHONY: build
build:
	go run cmd/api-server/main.go

test: 
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build