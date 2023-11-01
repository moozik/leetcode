package p849

import (
	"moozik/leetcode/utils"
	"testing"
)

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{seats: utils.IntArr("[1,0,0,1]")},
			want: 1,
		},
		{
			args: args{seats: utils.IntArr("[0,0,1]")},
			want: 2,
		},
		{
			args: args{seats: utils.IntArr("[0,1]")},
			want: 1,
		},
		{
			args: args{seats: utils.IntArr("[1,0,0,0,1,0,1]")},
			want: 2,
		},
		{
			args: args{seats: utils.IntArr("[1,0,0,0]")},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
