package tpl

// MakefileTemplate define a Makefile template
func MakefileTemplate() []byte {
	return []byte(`.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## clean artifacts
	@rm -rf bin charts coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: update-package
update-package: ## update package and commit
	@go get -u ./...
	@go mod tidy
	@git add go.mod go.sum
	@git commit -m "build: update package"

.PHONY: lint
lint: ## execute golint
	@golangci-lint run ./... -c .golangci.yaml

.PHONY: test-unit
test-unit: ## execute unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: gen
gen: gen-wire gen-mocks ## generate all generate commands

.PHONY: gen-wire
gen-wire: ## generate wire code
	@wire gen ./...

.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@protoc --proto_path=./pb --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./pb ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-mocks
gen-mocks: ## generate mocks code via mockery
	@go generate -tags=wireinject -x ./...
`)
}
