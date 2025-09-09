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
	@echo "Stopping Go container..."
	@docker compose down
	@echo "Go container stopped."

up:
	@echo "Starting Go container..."
	@docker compose up -d
	@echo "Go container started."

restart: down up
	@echo "Go container restarted."

logs:
	@echo "Fetching Go container logs..."
	@docker compose logs -f
	@echo "End of Go container logs."

shell:
	@docker compose exec ${SERVICE_NAME} bash

exit: down
	@echo "Removing Go container and image..."
	@docker image rm -f ${IMAGE_NAME} || true
	@echo "Go container and image removed."

.PHONY: launch build rebuild down up restart logs
