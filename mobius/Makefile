GO=go

default: all
all: build

vet:
	$(GO) vet ./...

build: gen vet
	scripts/build.sh
.PHONY: build

dbuild: gen vet
	GOOS=linux scripts/build.sh
.PHONY: dbuild

gen:
	sh -c "cd ../protos && make all"
.PHONY: gen

clean:
	rm -rf bin
.PHONY: clean

install-tools: install-go-tools
.PHONY: install-tools

install-go-tools:
	$(GO) install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
.PHONY: install-go-tools

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*?## "} /^[\/.a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help