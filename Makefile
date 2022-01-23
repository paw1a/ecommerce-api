build:
	go run cmd/app/main.go

build_swag:
	swag fmt -g cmd/app/main.go
	swag init -g cmd/app/main.go
	go run cmd/app/main.go

.DEFAULT_GOAL := build_swag
