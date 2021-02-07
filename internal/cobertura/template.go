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
# Test coverage

**Total:** {{.LineRate | humanizePercentage}}%</br>

**Lines Covered:** {{.LinesCovered}}</br>
**Lines Valid:** {{.LinesValid}}</br>
`

// RunTemplate run the report template
func RunTemplate(wr io.Writer, cr *CoverageReport) error {
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"humanizePercentage": humanizePercentage,
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
