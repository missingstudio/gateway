default: build

DEV_COMPOSE_FILE = docker-compose.yml

build: build-protos build-mobius
.PHONY: build

build-mobius:
	sh -c "cd ./mobius && make all"
.PHONY: build-mobius

build-protos:
	sh -c "cd ./protos && make all"
.PHONY: build-protos

clean: clean-backend clean-worker clean-cli
.PHONY: clean

clean-backend:
	sh -c "cd ./backend && make clean"
.PHONY: clean-backend

rebuild: clean build
.PHONY: rebuild

compose-dev: rebuild
	sh -c "docker compose -f $(DEV_COMPOSE_FILE) up -d"
.PHONY: compose-up