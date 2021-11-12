package p375

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			args:    args{n: 10},
			wantAns: 16,
		},
		{
			args:    args{n: 2},
			wantAns: 1,
		},
		{
			args:    args{n: 200},
			wantAns: 952,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getMoneyAmount(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("getMoneyAmount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
