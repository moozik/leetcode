package utils

import (
	"reflect"
	"testing"
)

func Test_genIntArray(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{str: "[[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]"},
			want: [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenIntArray(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
