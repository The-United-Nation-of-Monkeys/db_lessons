.PHONY: build up down restart logs clean help seed migrate-up migrate-down

# Default target
help:
	@echo "Available commands:"
	@echo "  make build        - Build Docker images"
	@echo "  make up           - Start all services"
	@echo "  make down         - Stop all services"
	@echo "  make restart      - Restart all services"
	@echo "  make logs         - Show logs from all services"
	@echo "  make clean        - Stop services and remove volumes"
	@echo "  make swagger      - Generate Swagger documentation"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make migrate-down - Rollback database migrations"
	@echo "  make seed         - Seed database with test data"

build:
	docker-compose build

up:
	docker-compose --env-file=config.env up -d

down:
	docker-compose down

restart:
	docker-compose restart

logs:
	docker-compose logs -f

clean:
	docker-compose down -v

swagger:
	@echo "Generating Swagger documentation..."
	@if ! command -v swag > /dev/null 2>&1; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@GOPATH=$$(go env GOPATH); \
	if [ -f "$$GOPATH/bin/swag" ]; then \
		$$GOPATH/bin/swag init -g cmd/app/main.go -o docs --parseDependency --parseInternal --exclude test-task-wallet; \
	elif [ -f "$$HOME/go/bin/swag" ]; then \
		$$HOME/go/bin/swag init -g cmd/app/main.go -o docs --parseDependency --parseInternal --exclude test-task-wallet; \
	else \
		swag init -g cmd/app/main.go -o docs --parseDependency --parseInternal --exclude test-task-wallet; \
	fi
	@echo "Swagger documentation generated successfully!"

migrate-up:
	@echo "Running database migrations..."
	@if ! command -v migrate > /dev/null 2>&1; then \
		echo "Installing golang-migrate..."; \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
	fi
	@GOPATH=$$(go env GOPATH); \
	DB_URL=$${DB_URL:-postgres://postgres:postgres@localhost:5432/online_classes?sslmode=disable}; \
	if [ -f "$$GOPATH/bin/migrate" ]; then \
		$$GOPATH/bin/migrate -path migrations -database "$$DB_URL" up; \
	elif [ -f "$$HOME/go/bin/migrate" ]; then \
		$$HOME/go/bin/migrate -path migrations -database "$$DB_URL" up; \
	else \
		migrate -path migrations -database "$$DB_URL" up; \
	fi
	@echo "Migrations completed!"

migrate-down:
	@echo "Rolling back database migrations..."
	@if ! command -v migrate > /dev/null 2>&1; then \
		echo "Installing golang-migrate..."; \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
	fi
	@GOPATH=$$(go env GOPATH); \
	DB_URL=$${DB_URL:-postgres://postgres:postgres@localhost:5432/online_classes?sslmode=disable}; \
	if [ -f "$$GOPATH/bin/migrate" ]; then \
		$$GOPATH/bin/migrate -path migrations -database "$$DB_URL" down; \
	elif [ -f "$$HOME/go/bin/migrate" ]; then \
		$$HOME/go/bin/migrate -path migrations -database "$$DB_URL" down; \
	else \
		migrate -path migrations -database "$$DB_URL" down; \
	fi
	@echo "Migrations rolled back!"

seed:
	@echo "Seeding database with test data..."
	@go run cmd/seed/main.go -file scripts/seed_data.sql
	@echo "Database seeded successfully!"

