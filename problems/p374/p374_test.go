package p374

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 100},
			want: 93,
		},
		{
			args: args{n: 50},
			want: 43,
		},
		{
			args: args{n: 2147483647},
			want: 857093,
		},
		{
			args: args{n: 1},
			want: 1,
		},
		{
			args: args{n: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
