package container

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a *[4]float64
	}
	tests := []struct {
		name    string
		args    args
		wantSum float64
	}{
		{"", args{a: &[4]float64{7.0, 8.5, 9.1, 6.0}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.a); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
