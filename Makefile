current_dir = $(shell pwd)
USER := $(shell id -u -n)
GIT_ROOT := $(shell git rev-parse --show-toplevel)
ARCH := $(shell uname -m)
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

trivy_version := 0.26.0
golangci_lint_version := v1.45.2
tag := latest
api_key := test

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s '>' | sed -e 's/^/ /'

.PHONY: install/lint
## install/lint > Install golangci-lint
install/lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin $(golangci_lint_version)
	golangci-lint version

.PHONY: imports
## imports > Format and fix all go imports errors
imports:
	@echo Running goimports...
	goimports -w -local github.com/dimaunx/go-clean-example $(shell find $(GIT_ROOT) -type f -name '*.go' -not -path "$(GIT_ROOT)/vendor/*")

.PHONY: vendor
## vendor > Run go mod vendor
vendor:
	@echo Running vendor...
	cd $(GIT_ROOT) && go mod tidy && go mod vendor

.PHONY: lint
## lint > Run golangci-lint
lint: vendor
	@echo Running linter...
	golangci-lint run -v --timeout=15m

.PHONY: build
## build > build docker image, make build tag=test if tag is omitted latest is used
build: vendor
	docker build . -t go-clean-example:$(tag)
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/trivycache:/root/.cache/ aquasec/trivy:$(trivy_version) image --exit-code 0 --severity HIGH,CRITICAL go-clean-example:$(tag)

.PHONY: run
## run > run server locally, make run tag=test if tag is omitted latest is used
run: build
	docker-compose up -d
	docker run --rm --network go-clean-example_default -p 8000:8000 -e REDIS_HOST=redis:6379 -e API_KEY=$(api_key) go-clean-example:$(tag)

.PHONY: push
## push > push docker image, make push tag=test if tag is omitted latest is used
push: build
	docker push go-clean-example:$(tag)