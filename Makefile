.DEFAULT_GOAL := help
.PHONY: help

VERSION=$(shell cat ./VERSION)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	cd cmd/adampresleydev && CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${VERSION}'" -mod=mod -o adampresleydev .

run: ## Run the application using mise and process-compose
	mise run run
