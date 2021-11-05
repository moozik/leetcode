package p1218

import (
	"testing"
)

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			want: 4,
		},
		{
			args: args{
				arr:        []int{1, 3, 5, 7},
				difference: 1,
			},
			want: 1,
		},
		{
			args: args{
				arr:        []int{1, 5, 7, 8, 5, 3, 4, 2, 1},
				difference: -2,
			},
			want: 4,
		},
		{
			args: args{
				arr:        []int{3,4,-3,-2,-4},
				difference: -5,
			},
			want: 2,
		},
		{
			args: args{
				arr:        []int{4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8},
				difference: 0,
			},
			want: 2,
		},
		{
			args: args{
				arr:        []int{16,-4,-6,-11,-8,-9,4,-11,15,15,-9,11,7,-7,10,-16,4},
				difference: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
