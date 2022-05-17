package p506

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{score: []int{10, 3, 8, 9, 4}},
			want: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			args: args{score: []int{5, 4, 3, 2, 1}},
			want: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			args: args{score: []int{5, 4, 3, 2, 1, 500, 100, 58, 6, 24, 11, 88, 499}},
			want: []string{"9", "10", "11", "12", "13", "Gold Medal", "Bronze Medal", "5", "8", "6", "7", "4", "Silver Medal"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
