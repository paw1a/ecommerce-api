build:
	go run cmd/app/main.go

build_swag:
	swag fmt -g cmd/app/main.go
	swag init -g cmd/app/main.go
	go run cmd/app/main.go

init_db:
	@cd ./db \
	&& pip3 install -r requirements.txt \
	&& python3 generate.py \
	&& (bash init.sh "ecommerce");

.DEFAULT_GOAL := build_swag
