package p319

import "testing"

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 1},
			want: 1,
		},
		{
			args: args{n: 2},
			want: 1,
		},
		{
			args: args{n: 3},
			want: 1,
		},
		{
			args: args{n: 4},
			want: 2,
		},
		{
			args: args{n: 5},
			want: 2,
		},
		{
			args: args{n: 6},
			want: 2,
		},
		{
			args: args{n: 20},
			want: 4,
		},
		{
			args: args{n: 301},
			want: 17,
		},
		{
			args: args{n: 1000000000},
			want: 31622,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
