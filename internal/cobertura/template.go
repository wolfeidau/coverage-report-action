package cobertura

import (
	"fmt"
	"io"
	"strconv"
	"text/template"

	"github.com/dustin/go-humanize"
	"github.com/rs/zerolog/log"
)

var reportTemplate = `
![coverage](https://img.shields.io/badge/coverage%20total-{{.LineRate | humanizePercentage}}%25-{{.LineRate | badgeColor}}?style=for-the-badge)

| Package  | Coverage of Statements | threshold ({{threshold}}%) |
| ------------- | ------------- | ------------- |
{{- range .Packages.Package }}
| {{ .Name }}  | {{.LineRate | humanizePercentage}}% | {{.LineRate | thresholdEmoji}} |
{{- end }}
| **total** | {{.LineRate | humanizePercentage}}% | {{.LineRate | thresholdEmoji}} |
`

// RunTemplate run the report template
func RunTemplate(wr io.Writer, cr *CoverageReport, minimumCoverage int) error {
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"humanizePercentage": humanizePercentage,
		"badgeColor":         thresholdString(minimumCoverage, "green", "red"),
		"thresholdEmoji":     thresholdString(minimumCoverage, "✅", "❌"),
		"threshold":          templateInt(minimumCoverage),
	}

	tmpl, err := template.New("test").Funcs(funcMap).Parse(reportTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse report template: %w", err)
	}

	err = tmpl.Execute(wr, cr)
	if err != nil {
		return fmt.Errorf("failed to execute report template: %w", err)
	}

	return nil
}

func humanizePercentage(val string) string {
	ft, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Fatal().Msg("failed to read float")
	}

	return humanize.FormatFloat("###.##", ft*100)
}

func thresholdString(minCoverage int, success, failure string) func(val string) string {
	return func(val string) string {

		ft, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal().Msg("failed to read float")
		}

		fmt.Println(minCoverage, ft)

		if int(ft*100) < minCoverage {
			return failure
		}

		return success
	}
}

func templateInt(val int) func() string {
	return func() string {
		return fmt.Sprint(val)
	}
}
