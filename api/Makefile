.DEFAULT_GOAL := build

.PHONY: fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o mdg-inventory-api

clean:
	go clean