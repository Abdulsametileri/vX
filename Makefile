## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ":" | sed -e 's/^/  /'

## lint: runs golangci lint based on .golangci.yml configuration
.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml  --fix -v

## unit-test: runs tests
.PHONY: unit-test
unit-test:
	go test -v ./... -coverprofile=unit_coverage.out

## unit-coverage-html: extract unit tests coverage to html format
.PHONY: unit-coverage-html
unit-coverage-html:
	make unit-test
	go tool cover -html=unit_coverage.out -o unit_coverage.html

## build-cli: build the cli application
.PHONY: build-cli
build-cli:
	go build -o vx cli/main.go