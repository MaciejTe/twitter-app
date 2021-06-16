CWD=$$(pwd)
PKG := "/app"
PKG_LIST := $(shell go list ${PKG}/...)

.PHONY: all dep build test coverage coverhtml lint

all: build

build_image:
	docker build -t twitter_app .

build_image_dev:
	docker build -t twitter_app_dev -f Dockerfile.dev .

build_image_tests:
	docker build -t twitter_app_tests -f tests/Dockerfile tests/

dev-tests: ## Start Python tests container
	docker run --network host --rm -it --env-file .env -v "${CWD}/tests":/tests -w /tests twitter_app_tests bash

run:
	docker run --network host --rm -it --env-file .env twitter_app

dev:
	docker run --network host --rm -it --env-file .env -v "${CWD}":/app/ twitter_app_dev bash


# FOLLOWING COMMANDS HAVE TO BE RUN INSIDE TWITTER DOCKER CONTAINER

lint: ## Lint the files
	@make dep
	@gofmt -w pkg/ api/ main.go
	@goimports -w pkg/ api/ main.go
	@go vet .
	@golint -set_exit_status ${PKG_LIST}
	@goimports -w pkg/ api/ main.go
	@gocyclo -over 15 pkg/ api/ main.go
	@golangci-lint run
	@go build -o twitter_app; rm -f twitter_app
	@go mod tidy

coverage: ## Generate global code coverage report ()
	@go test -covermode=atomic -coverprofile coverage.out -v ./...

coverhtml: coverage ## Generate global code coverage report in HTML
	@go tool cover -html=coverage.out -o coverage.html


dep: ## Get the test/lint dependencies
	@go get -u golang.org/x/lint/golint
	@go get github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	@go get golang.org/x/tools/cmd/goimports

build: ## Build Twitter application
	go build -o twitter_app

test: ## Run all go-based tests
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
