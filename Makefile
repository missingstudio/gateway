default: build

DEV_COMPOSE_FILE = docker-compose.yml

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*?## "} /^[\/.a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

build: build-mobius ## Build missing studio project
.PHONY: build

build-mobius:
	sh -c "cd ./mobius && make all"
.PHONY: build-mobius

build-protos: ## Build missing studio protos
	sh -c "cd ./protos && make all"
.PHONY: build-protos

clean: clean-mobius ## Clean missing studio
.PHONY: clean

clean-mobius:
	@echo "🧹 Cleaning mobius.."
	sh -c "cd ./mobius && make clean"
.PHONY: clean-mobius

rebuild: clean build ## Rebuild missing studio again
.PHONY: rebuild

compose-dev: rebuild ## Start a missing studio container
	sh -c "docker compose -f $(DEV_COMPOSE_FILE) up -d"
.PHONY: compose-up