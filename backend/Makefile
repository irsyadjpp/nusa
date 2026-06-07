.PHONY: help build run test lint migrate-up migrate-down migrate-create clean install-deps docker-up docker-down generate seed

help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make test           - Run tests"
	@echo "  make lint           - Run linter"
	@echo "  make migrate-up     - Run database migrations up"
	@echo "  make migrate-down   - Run database migrations down"
	@echo "  make migrate-create - Create a new migration (usage: make migrate-create NAME=migration_name)"
	@echo "  make generate       - Generate code from sqlc"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make install-deps   - Install dependencies"
	@echo "  make docker-up      - Start Docker Compose services"
	@echo "  make docker-down    - Stop Docker Compose services"
	@echo "  make seed           - Seed database with initial data"

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...

migrate-up:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)

generate:
	sqlc generate

clean:
	rm -rf bin/

install-deps:
	go mod download
	go mod tidy

docker-up:
	docker compose up -d

docker-down:
	docker compose down

seed:
	go run cmd/seed/main.go
