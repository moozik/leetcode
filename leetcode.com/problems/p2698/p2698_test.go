package p2698

import "testing"

func Test_dp(t *testing.T) {
	type args struct {
		s      string
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s:      "81",
				target: 9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("dp() = %v, want %v", got, tt.want)
			}
		})
	}
}
