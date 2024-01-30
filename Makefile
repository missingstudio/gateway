default: build

DEV_COMPOSE_FILE = docker-compose.yml

build: build-mobius
.PHONY: build

build-mobius:
	sh -c "cd ./mobius && make all"
.PHONY: build-mobius

build-protos:
	sh -c "cd ./protos && make all"
.PHONY: build-protos

clean: clean-mobius
.PHONY: clean

clean-mobius:
	sh -c "cd ./mobius && make clean"
.PHONY: clean-mobius

rebuild: clean build
.PHONY: rebuild

compose-dev: rebuild
	sh -c "docker compose -f $(DEV_COMPOSE_FILE) up -d"
.PHONY: compose-up