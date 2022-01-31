build:
	go run cmd/app/main.go

swag:
	swag fmt -g cmd/app/main.go
	swag init -g cmd/app/main.go
	go run cmd/app/main.go

init:
	@cd ./db \
	&& pip3 install -r requirements.txt \
	&& python3 generate.py \
	&& (bash init.sh "ecommerce");

.DEFAULT_GOAL := build
