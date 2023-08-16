SHELL := /bin/bash
.DEFAULT_GOAL := help

###########################
# VARIABLES
###########################
VERSION := v12
COMMIT_SHA :=
COMMIT_DATE :=
LD_FLAGS := -s -w -X github.com/AndreasAugustin/go-gitmoji-cli/pkg.Version=${VERSION} -X github.com/AndreasAugustin/go-gitmoji-cli/pkg.CommitSHA=${COMMIT_SHA} -X github.com/AndreasAugustin/go-gitmoji-cli/pkg.CommitDate=${COMMIT_DATE}

###########################
# MAPPINGS
###########################

###########################
# TARGETS
###########################

.PHONY: help
help:  ## help target to show available commands with information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: markdownlint
markdownlint: ## Validate markdown files
	docker-compose run docs markdownlint .github/ --ignore node_modules
	docker-compose run docs markdownlint . --ignore node_modules

.PHONY: golangci-lint
golangci-lint:  ## run golangci-lint https://golangci-lint.run/
	docker-compose run golangci-lint golangci-lint run -v

.PHONY: zsh
zsh: ## open dev container with build environment
	docker-compose run --service-ports dev /bin/zsh

.PHONY: prune
prune: ## delete the whole environment
	docker-compose down -v --rmi all --remove-orphans

.PHONY: build
build: ## build the solution
	go build -ldflags="${LD_FLAGS}" -o out/

.PHONY: test
test:  ## run the tests
	go test -ldflags="${LD_FLAGS}" ./... -cover

.PHONY: lint
lint:
	go vet ./...

.PHONY: format
format:  ## format the files
	go fmt ./...
	go fix ./...

.PHONY: clean
clean:  ## clean
	go clean
	go mod tidy

.PHONY: install
install: ## install the package
	go install

.PHONY: create-gifs
create-gifs: ## create the gifs
	$(MAKE) -C docs/ vhs
