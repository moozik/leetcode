package main

import "testing"

func Test_isVaild(t *testing.T) {
	type args struct {
		num int
		x   int
		y   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				num: 0,
				x:   4,
				y:   4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVaild(tt.args.num, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isVaild() = %v, want %v", got, tt.want)
			}
		})
	}
}
