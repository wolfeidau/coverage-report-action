package cobertura

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_humanizePercentage(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should render percentage",
			args: args{val: ".01"},
			want: "1.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := humanizePercentage(tt.args.val); got != tt.want {
				t.Errorf("humanizePercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateColor(t *testing.T) {
	type args struct {
		minimumCoverage int
		val             string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should render percentage",
			args: args{val: ".01", minimumCoverage: 10},
			want: "red",
		},
		{
			name: "should render percentage",
			args: args{val: "20.01", minimumCoverage: 10},
			want: "green",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thresholdString(tt.args.minimumCoverage, "green", "red")(tt.args.val); got != tt.want {
				t.Errorf("generateColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunTemplate(t *testing.T) {
	assert := require.New(t)
	type args struct {
		fixture         string
		minimumCoverage int
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		{
			name: "should render report",
			args: args{fixture: "fixtures/coverage.xml", minimumCoverage: 10},
			wantWr: `
![coverage](https://img.shields.io/badge/coverage%20total-11.65%25-green?style=for-the-badge)

| Package  | Coverage of Statements | threshold (10%) |
| ------------- | ------------- | ------------- |
| github.com/wolfeidau/coverage-report-action/internal/cobertura  | 38.71% | ✅ |
| github.com/wolfeidau/coverage-report-action/internal/flags  | 0.00% | ❌ |
| github.com/wolfeidau/coverage-report-action  | 0.00% | ❌ |
| **total** | 11.65% | ✅ |
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}

			cr := new(CoverageReport)
			data, err := ioutil.ReadFile(tt.args.fixture)
			assert.NoError(err)
			err = xml.NewDecoder(bytes.NewBuffer(data)).Decode(cr)
			assert.NoError(err)

			if err := RunTemplate(wr, cr, tt.args.minimumCoverage); (err != nil) != tt.wantErr {
				t.Errorf("RunTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t.Errorf("RunTemplate() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
