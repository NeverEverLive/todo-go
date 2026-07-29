include .env
export

export PROJECT_ROOT := $(shell pwd)
export CURRENT_USER_HOST := $(shell id -u):$(shell id -g)

env-up:
	@docker compose up todo-app-postgres -d

env-down:
	@docker compose down todo-app-postgres

env-drop:
	@read -p "Are you sure you want to drop the database? (y/n): " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todo-app-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Database dropped"; \
	else \
		echo "Operation aborted"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-forward-stop:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Please provide a sequence number. Ex: make migrate-create seq=1"; \
		exit 1; \
	fi; \
	docker compose run --rm todo-app-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Please provide an action. Ex: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm todo-app-postgres-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todo-app-postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		"$(action)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

todo-app-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go

logs-cleanup:
	@read -p "Are you sure you want to clean all logs? (y/n): " ans; \
    	if [ "$$ans" = "y" ]; then \
    		rm -rf ${PROJECT_ROOT}/out/logs && \
    		echo "Logs cleaned"; \
    	else \
    		echo "Operation aborted"; \
    	fi
