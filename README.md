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

## `minimum-coverage`

**Required** Minimum allowed coverage as an integer, eg. `80` represents 80% minimum lines covered.

## `show-files`

Toggles file results in the report output, defaults to `false`.

# Usage

```yaml
# must use pull request to ensure all required attributes are present for the coverage comment 
# action to be able to create/update the comment
on: [pull_request]
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go
        uses: actions/setup-go@v1
        with:
          go-version: 1.15.x
        id: go	
      - name: Check out code into the Go module directory
        uses: actions/checkout@v1
	  # run make ci target which runs tests and outputs coverage.xml
      - name: CI Tasks
        run: make ci
	  # run the report task which loads coverage.xml and then posts a comment with coverage summary
	  - name: Report
        uses: wolfeidau/coverage-report-action@v1
          with:
          github-token: ${{ secrets.github_token }}
          coverage-report: coverage.xml
          minimum-coverage: 80
```

# Output

This will maintain a single comment in your PR, which is updated each time changes are pushed to the pull request.

![coverage](https://img.shields.io/badge/coverage%20total-11.65%25-green?style=for-the-badge)

| Package/File | Coverage of Lines | Threshold (10%) |
| ------------- | ------------- | ------------- |
| **Total** | 11.65% | ✅ |
| internal/cobertura/cobertura.go | 0.00% | ❌ |
| internal/cobertura/template.go | 57.14% | ✅ |
| internal/flags/flags.go | 0.00% | ❌ |
| main.go | 0.00% | ❌ |

Generated by https://github.com/wolfeidau/coverage-report-action

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).