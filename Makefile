APP_NAME := $(shell basename $(CURDIR))

run:
	go run .

test:
	go test -v ./...

build:
	go build -o bin/$(APP_NAME) .

.PHONY: all test clean