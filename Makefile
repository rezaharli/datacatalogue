PROJECT_NAME := "datacatalogue"
PKG := "eaciit/$(PROJECT_NAME)/webapp"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test race lint

all: build

lint: ## Lint the files
	@cd webapp; golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@cd webapp; go test -short ./...

race: dep ## Run data race detector
	@cd webapp; go test -race -short ./...

build: dep ## Build the binary file
	@go build -o $(PROJECT_NAME) -i -v $(PKG) 

clean: cd ## Remove previous build
	@rm -f $(PROJECT_NAME)

dep: ## Get the dependencies
	@cd webapp; go get -v -d -t ./...
	@go get -u github.com/golang/lint/golint

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
