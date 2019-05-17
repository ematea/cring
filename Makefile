NAME := cring
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Install required tools
setup:
	go get github.com/golang/dep/cmd/dep
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

## Install dependencies
deps: setup
	dep ensure

## Update dependencies
update: setup
	dep ensure -update

## Lint
lint: setup
	go vet $$(go list ./...)
	for pkg in $$(go list ./...); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

## Format source codes
fmt: setup
	goimports -w -d $$(find . -type f -name '*.go' -not -path "./vendor/*")

## Run test
test: deps
	go test -v $$(go list ./...)

## Build cring command
build: deps
	go build -ldflags "$(LDFLAGS)" -o bin/cring cmd/cring/main.go

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.DEFAULT_GOAL := build
.PHONY: setup deps test update lint build help

