package cobertura

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
)

// CoverageReport generated from the coverage-04.dtd
type CoverageReport struct {
	XMLName         xml.Name `xml:"coverage"`
	Text            string   `xml:",chardata"`
	LinesValid      string   `xml:"lines-valid,attr"`
	LinesCovered    string   `xml:"lines-covered,attr"`
	LineRate        string   `xml:"line-rate,attr"`
	BranchesValid   string   `xml:"branches-valid,attr"`
	BranchesCovered string   `xml:"branches-covered,attr"`
	BranchRate      string   `xml:"branch-rate,attr"`
	Timestamp       string   `xml:"timestamp,attr"`
	Complexity      string   `xml:"complexity,attr"`
	Version         string   `xml:"version,attr"`
	Sources         struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source"`
	} `xml:"sources"`
	Packages struct {
		Text    string    `xml:",chardata"`
		Package []Package `xml:"package"`
	} `xml:"packages"`
}

// Package represents a package in cobertura format
type Package struct {
	Text       string `xml:",chardata"`
	Name       string `xml:"name,attr"`
	LineRate   string `xml:"line-rate,attr"`
	BranchRate string `xml:"branch-rate,attr"`
	Classes    struct {
		Text  string  `xml:",chardata"`
		Class []Class `xml:"class"`
	} `xml:"classes"`
}

// Class represents a Class or a File in cobertura format
type Class struct {
	Text       string `xml:",chardata"`
	Name       string `xml:"name,attr"`
	Filename   string `xml:"filename,attr"`
	LineRate   string `xml:"line-rate,attr"`
	BranchRate string `xml:"branch-rate,attr"`
	Methods    struct {
		Text   string `xml:",chardata"`
		Method []struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Hits      string `xml:"hits,attr"`
			Signature string `xml:"signature,attr"`
			Lines     struct {
				Text string `xml:",chardata"`
				Line struct {
					Text   string `xml:",chardata"`
					Number string `xml:"number,attr"`
					Hits   string `xml:"hits,attr"`
				} `xml:"line"`
			} `xml:"lines"`
		} `xml:"method"`
	} `xml:"methods"`
	Lines struct {
		Text string `xml:",chardata"`
		Line []struct {
			Text              string `xml:",chardata"`
			Number            string `xml:"number,attr"`
			Hits              string `xml:"hits,attr"`
			Branch            string `xml:"branch,attr"`
			ConditionCoverage string `xml:"condition-coverage,attr"`
		} `xml:"line"`
	} `xml:"lines"`
}

// ParseCoverageReport parse a cobertura coverage report
func ParseCoverageReport(r io.Reader) (*CoverageReport, error) {
	dec := xml.NewDecoder(r)

	report := new(CoverageReport)

	err := dec.Decode(report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode coverage report: %w", err)
	}

	return report, nil
}

// MeetsThreshold check if the line rate meets the minimum coverage threshold
func MeetsThreshold(lineRate string, minimumCoverage int) (bool, error) {
	ft, err := strconv.ParseFloat(lineRate, 64)
	if err != nil {
		return false, err
	}

	return int(ft*100) > minimumCoverage, nil
}
