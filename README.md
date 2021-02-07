# coverage-report-action

This action provides a summary comment for PRs which produce a cobertura format coverage file.

# Inputs

## `coverage-report`

**Required** The path to the cobertura format coverage file.

# Usage

```
uses: wolfeidau/coverage-report-action@v1
with:
    coverage-report: coverage.xml
```

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).