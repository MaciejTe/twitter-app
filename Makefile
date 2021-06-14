CWD=$$(pwd)
PKG := "/app"
PKG_LIST := $(shell go list ${PKG}/...)

.PHONY: all dep build test coverage coverhtml lint

all: build

build_image:
	docker build -t twitter_app .

build_image_dev:
	docker build -t twitter_app_dev -f Dockerfile.dev .

run:
	docker run --network host --rm -it --env-file .env twitter_app

dev:
	docker run --network host --rm -it --env-file .env -v "${CWD}":/app/ twitter_app_dev bash

build:
	go build -o twitter_app

test:
	go test

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

coverage: ## Generate global code coverage report ()
	@go test -covermode=atomic -coverprofile coverage.out -v ./...

coverhtml: coverage ## Generate global code coverage report in HTML
	@go tool cover -html=coverage.out -o coverage.html


dep: ## Get the dependencies
	@go get -u golang.org/x/lint/golint

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
