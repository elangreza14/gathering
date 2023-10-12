#!make
include local.env

up-stack:
	docker-compose --env-file local.env up --build -d

up:
	docker-compose --env-file local.env up -d

down:
	docker-compose --env-file local.env down

run-local:
	MODE=local go run cmd/server/main.go
	
run-live:
	go run cmd/server/main.go

test:
	 go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out
	
lint:
	gofumpt -l -w .

mock:
	go generate ./...

FILENAME?=file-name

migrate-down:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}" -path ${MIGRATION_FOLDER} down
	
migrate-create:
	@read -p  "What is the name of migration?" NAME; \
	migrate create -ext sql -tz Asia/Jakarta -dir ${MIGRATION_FOLDER} -format "20060102150405" $$NAME

.PHONY: up-stack up down run-local run-live mock lint test migrate-down migrate-create