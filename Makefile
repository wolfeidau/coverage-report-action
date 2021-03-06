GOLANGCI_VERSION = 1.32.0

ci: clean awsmocks lint test
.PHONY: ci

LDFLAGS := -ldflags="-s -w"

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

bin/gocover-cobertura:
	@go get -u github.com/boumenot/gocover-cobertura
	@env GOBIN=$$PWD/bin GO111MODULE=on go install github.com/boumenot/gocover-cobertura

clean:
	@echo "--- clean all the things"
	@rm -rf dist
.PHONY: clean

lint: bin/golangci-lint
	@echo "--- lint all the things"
	@bin/golangci-lint run
.PHONY: lint

test: bin/gocover-cobertura
	@echo "--- test all the things"
	@go test -coverprofile=coverage.txt -covermode count ./...
	@bin/gocover-cobertura < coverage.txt > coverage.xml
.PHONY: test
