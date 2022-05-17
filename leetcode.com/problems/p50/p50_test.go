package p50

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns float64
	}{
		{
			args:    args{x: 2, n: 10},
			wantAns: 1024,
		},
		// {
		// 	args:    args{x: 2.10000, n: 3},
		// 	wantAns: 9.261,
		// },
		// {
		// 	args:    args{x: 2, n: -2},
		// 	wantAns: 0.25,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := myPow(tt.args.x, tt.args.n); gotAns != tt.wantAns {
				t.Errorf("myPow() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
