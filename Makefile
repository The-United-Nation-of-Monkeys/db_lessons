.PHONY: build up down restart logs clean help

# Default target
help:
	@echo "Available commands:"
	@echo "  make build    - Build Docker images"
	@echo "  make up       - Start all services"
	@echo "  make down     - Stop all services"
	@echo "  make restart  - Restart all services"
	@echo "  make logs     - Show logs from all services"
	@echo "  make clean    - Stop services and remove volumes"
	@echo "  make swagger  - Generate Swagger documentation"

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

