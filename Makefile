.PHONY: build
build:
	go build -v ./cmd/server

.DEFAULT_GOAL := build
