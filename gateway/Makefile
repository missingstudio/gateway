GO=go

default: all
all: build

vet:
	@echo "🧹 Vetting go.mod.."
	$(GO) vet ./...

lint: ## Run linters
	golangci-lint run

lint-fix:
	golangci-lint run --fix

build: vet ## Build AI Gateway server
	scripts/build.sh
.PHONY: build

dbuild: vet
	GOOS=linux scripts/build.sh
.PHONY: dbuild

clean: ## Clean AI Gateway server
	rm -rf bin
.PHONY: clean

test: ## Run tests
	@go test -race $(shell go list ./... | grep -v /ui | grep -v /vendor/ | grep -v /test/) -coverprofile=coverage.out -count 2 -timeout 150s

install-tools: install-go-tools
.PHONY: install-tools

install-go-tools:
	$(GO) install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
.PHONY: install-go-tools

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*?## "} /^[\/.a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help