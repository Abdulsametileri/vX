lint:
	golangci-lint run -c .golangci.yml  --fix -v

unit-test:
	go test -v ./... -coverprofile=unit_coverage.out

unit-coverage-html:
	make unit-test
	go tool cover -html=unit_coverage.out -o unit_coverage.html

build-cli:
	go build -o vx cli/main.go