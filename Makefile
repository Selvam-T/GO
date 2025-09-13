# Makefile for managing Go development environment with Docker
include .env
export

launch:	build up shell
go:	build up shell

build:
	@echo "Building Go image..."
	@docker compose build --no-cache
	@echo "Go image built successfully."

rebuild: down build up
	@echo "Go container rebuilt and started."

down:
	@echo "Stopping Go containers..."
	@docker compose down
	@echo "Go containers stopped."

up:
	@echo "Starting Containers..."
	#@docker compose up -d
	@docker-compose up -d
	@echo "Container started."

restart: down up
	@echo "Containers restarted."

logs:	
	@echo "Fetching all container logs in the Compose project ..."
	@docker compose logs -f
	@echo "End of Go container logs."

logs-db:
	@echo "Fetching Postgres logs..."
	@docker compose logs -f ${DB_SERVICE}

logs-go:
	@echo "Fetching Go service logs..."
	@docker compose logs -f ${GO_SERVICE}

shell:
	@echo "+-----------------------------------+"
	@echo "|              |                    |"
	@echo "|              v                    |"
	@echo "| You are inside the container now. |"
	@echo "+-----------------------------------+"
	@docker compose exec ${GO_SERVICE} bash

exit: down
	@echo "Removing built images..."
	@docker image rm -f ${GO_IMAGE_NAME} || true
	@docker image rm -f ${DB_IMAGE_NAME} || true
	@echo "Image removed."
	
prune:	exit
	@echo "Cleaning up Docker build cache..."
	@docker system prune -af || true
	@docker volume prune -f
	@echo "System and volume pruning done."

.PHONY: launch go build rebuild down up restart logs logs-db logs-go shell exit prune
