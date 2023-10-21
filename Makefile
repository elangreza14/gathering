GO ?= go
GOBIN ?= $$($(GO) env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOLANG_TEST_COVERAGE ?= $(GOBIN)/golang-test-coverage
GOLANG_SWAG ?= $(GOBIN)/swag
GOLANGCI_LINT_VERSION ?= v1.52.2

#!make
include local.env

up-stack:
	docker-compose --env-file local.env up --build -d

up:
	docker-compose --env-file local.env up -d

down:
	docker-compose --env-file local.env down

install-swag:
	test -f $(GOLANG_SWAG) || go install github.com/swaggo/swag/cmd/swag@latest

run-local: install-swag
	swag init -g cmd/server/main.go
	swag fmt
	MODE=local go run cmd/server/main.go
	
run-live:
	go run cmd/server/main.go

get-golangcilint:
	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

lint: get-golangcilint
	$(GOLANGCI_LINT) run ./...

test: 
	go test -timeout=3s -race -count=10 -failfast -short ./...
	go test -timeout=3s -race -count=1 -failfast ./...

tidy:
	go mod tidy
	go fmt ./...

install-go-test-coverage:
	test -f $(GOLANG_TEST_COVERAGE) || go install github.com/vladopajic/go-test-coverage/v2@latest

check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./coverage.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.github/.testcoverage.yml
	go tool cover -html=coverage.out

test-coverage:
	go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out

mock:
	go generate ./...

FILENAME?=file-name

migrate-down:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}" -path ${MIGRATION_FOLDER} down
	
migrate-create:
	@read -p  "What is the name of migration?" NAME; \
	migrate create -ext sql -tz Asia/Jakarta -dir ${MIGRATION_FOLDER} -format "20060102150405" $$NAME

.PHONY: get-golangcilint lint test tidy check-coverage up-stack up down run-local run-live mock migrate-down migrate-create install-go-test-coverage install-swag run-swag