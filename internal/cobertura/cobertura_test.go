package cobertura

import "testing"

func TestMeetsThreshold(t *testing.T) {
	type args struct {
		lineRate        string
		minimumCoverage int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "should return true for above threshold",
			args:    args{lineRate: "0.12", minimumCoverage: 10},
			want:    true,
			wantErr: false,
		},
		{
			name:    "should return false for below threshold",
			args:    args{lineRate: "0.12", minimumCoverage: 30},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MeetsThreshold(tt.args.lineRate, tt.args.minimumCoverage)
			if (err != nil) != tt.wantErr {
				t.Errorf("MeetsThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MeetsThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}
