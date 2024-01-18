default: build

build: build-protos build-backend
.PHONY: build

build-backend:
	sh -c "cd ./backend && make all"
.PHONY: build-backend

build-protos:
	sh -c "cd ./protos && make all"
.PHONY: build-protos