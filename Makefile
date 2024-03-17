default: build

DEV_COMPOSE_FILE = docker-compose.yml

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*?## "} /^[\/.a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

build: build-gateway ## Build missing studio project
.PHONY: build

test:
	sh -c "cd ./gateway && make test"
.PHONY: test

build-gateway:
	sh -c "cd ./gateway && make all"
.PHONY: build-gateway

build-protos: ## Build missing studio protos
	sh -c "cd ./protos && make all"
.PHONY: build-protos

clean: clean-gateway ## Clean missing studio
.PHONY: clean

clean-gateway:
	@echo "Cleaning gateway.."
	sh -c "cd ./gateway && make clean"
.PHONY: clean-gateway

rebuild: clean build ## Rebuild missing studio again
.PHONY: rebuild

compose-dev: rebuild ## Start a missing studio container
	sh -c "docker compose -f $(DEV_COMPOSE_FILE) up -d"
.PHONY: compose-up