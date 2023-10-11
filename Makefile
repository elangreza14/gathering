#!make
include local.env

up-stack:
	docker-compose --env-file local.env up --build -d

up:
	docker-compose --env-file local.env up -d

down:
	docker-compose --env-file local.env down

run-local:
	MODE=local go run main.go
	
run-live:
	go run main.go

test:
	 go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out
	
lint:
	gofumpt -l -w .

.PHONY: up-stack up down run-local run-live migrate-create mock lint test