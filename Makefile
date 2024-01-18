default: build

build: build-backend
.PHONY: build

build-backend:
	sh -c "cd ./backend && make all"
.PHONY: build-backend