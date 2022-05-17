package p397

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 16},
			want: 4,
		},
		{
			args: args{n: 15},
			want: 5,
		},
		{
			args: args{n: 13},
			want: 5,
		},
		{
			args: args{n: 12},
			want: 4,
		},
		{
			args: args{n: 11},
			want: 5,
		},
		{
			args: args{n: 10},
			want: 4,
		},
		{
			args: args{n: 9},
			want: 4,
		},
		{
			args: args{n: 8},
			want: 3,
		},
		{
			args: args{n: 7},
			want: 4,
		},
		{
			args: args{n: 6},
			want: 3,
		},
		{
			args: args{n: 5},
			want: 3,
		},
		{
			args: args{n: 4},
			want: 2,
		},
		{
			args: args{n: 3},
			want: 2,
		},
		{
			args: args{n: 2},
			want: 1,
		},
		{
			args: args{n: 1},
			want: 0,
		},
		{
			args: args{n: 2147483647},
			want: 32,
		},
		{
			args: args{n: 2147483646},
			want: 32,
		},
		{
			args: args{n: 2147483645},
			want: 33,
		},
		{
			args: args{n: 2147483644},
			want: 32,
		},
		{
			args: args{n: 2147483643},
			want: 33,
		},
		{
			args: args{n: 2147483642},
			want: 33,
		},
		{
			args: args{n: 2147483641},
			want: 33,
		},
		{
			args: args{n: 2147483640},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
