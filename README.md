# coverage-report-action

This action provides a summary comment for PRs which produce a [cobertura](https://cobertura.github.io/cobertura/) format file containing coverage. The aim is to provide an action for any language which supports outputting this report format.

# Why cobertura?

This XML format is supported across a range of different languages due to it's use in Jenkins, so it makes sense to use it for this action.

# Project Setup

## Go

To enable Go projects to produce a cobertura report file you will need to add the following to the makefile.

```Makefile
# install a local copy of a CLI tool which converts Go coverage to cobertura format.
bin/gocover-cobertura:
	@go get -u github.com/boumenot/gocover-cobertura
	@env GOBIN=$$PWD/bin GO111MODULE=on go install github.com/boumenot/gocover-cobertura

# run tests and capture coverage then generate the cobertura format report in xml
test: bin/gocover-cobertura
	@echo "--- test all the things"
	@go test -coverprofile=coverage.txt -covermode count ./...
	@bin/gocover-cobertura < coverage.txt > coverage.xml
.PHONY: test
```

# Inputs

## `github-token`

**Required** Github token for the repository.

## `coverage-report`

**Required** The path to the cobertura format coverage file.

# Usage

```
uses: wolfeidau/coverage-report-action@v1
with:
    github-token: ${{ secrets.github_token }}
    coverage-report: coverage.xml
```

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).