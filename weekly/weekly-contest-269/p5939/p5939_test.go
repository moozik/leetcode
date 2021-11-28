package p5939

import (
	"reflect"
	"testing"
)

func Test_getAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			args:    args{nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, k: 3},
			wantAns: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			args:    args{nums: []int{100000}, k: 0},
			wantAns: []int{100000},
		},
		{
			args:    args{nums: []int{8}, k: 100000},
			wantAns: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getAverages(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getAverages() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
