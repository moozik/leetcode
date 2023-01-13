package p1

import (
	"moozik/leetcode/utils"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:   utils.IntArr("[2,7,11,15]"),
				target: 9,
			},
			want: utils.IntArr("[0,1]"),
		},
		{
			args: args{
				nums:   utils.IntArr("[3,2,4]"),
				target: 6,
			},
			want: utils.IntArr("[1,2]"),
		},
		{
			args: args{
				nums:   utils.IntArr("[3,3]"),
				target: 6,
			},
			want: utils.IntArr("[0,1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
