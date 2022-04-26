include .env
export

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go && env

run: build
	docker-compose up app redis db

clean:
	rm -rf .bin .data

swag:
	swag fmt -g cmd/app/main.go
	swag init -g cmd/app/main.go
	go run cmd/app/main.go

init:
	@cd ./db \
	&& pip3 install -r requirements.txt \
	&& python3 generate.py \
	&& (bash init.sh "ecommerce");

test:
	go test -v ./internal/service/

.DEFAULT_GOAL := run
