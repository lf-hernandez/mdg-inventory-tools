.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o bin/mdg-inventory-api

run: build
	./bin/mdg-inventory-api

clean:
	go clean
