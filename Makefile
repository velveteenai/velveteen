export GO111MODULE := on

GOLANGCI_LINT_VERSION := v1.51.2
GIT_REVISION ?= latest
TEST_FILES ?= ./...
TEST_PATTERN ?= .
TEST_OPTIONS ?=

bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_LINT_VERSION)

test-results/go:
	mkdir -p test-results/go

setup:
	go mod download
	go install gotest.tools/gotestsum@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golang/mock/mockgen@latest
.PHONY: setup

build:
	CGO_ENABLED=0 go build -ldflags="-X main.AppRevision=$(GIT_REVISION)" -a -o ./velveteen ./cmd/http-server/.
.PHONY: build

test: test-results/go
	gotestsum --junitfile test-results/go/results.xml -- $(TEST_OPTIONS) -failfast -race -coverpkg=./... \
		-covermode=atomic -coverprofile=coverage.txt $(TEST_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

fmt:
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

lint: bin/golangci-lint
	bin/golangci-lint run ./...
.PHONY: lint

clean:
	go clean ./...
	rm -f ./velveteen
	rm -f ./coverage.txt
	rm -rf ./test-results
.PHONY: clean

.DEFAULT_GOAL := build
