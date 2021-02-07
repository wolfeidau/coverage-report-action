package cobertura

import (
	"bytes"
	"testing"
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

func TestRunTemplate(t *testing.T) {
	type args struct {
		cr *CoverageReport
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		{
			name: "should render report",
			args: args{&CoverageReport{
				LineRate:     "0.6526",
				LinesCovered: "325",
				LinesValid:   "498",
			}},
			wantWr: `
**coverage:** 65.26% of statements
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			if err := RunTemplate(wr, tt.args.cr); (err != nil) != tt.wantErr {
				t.Errorf("RunTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t.Errorf("RunTemplate() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
