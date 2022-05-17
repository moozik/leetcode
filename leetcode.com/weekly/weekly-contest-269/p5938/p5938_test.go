package p5938

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			args:    args{nums: []int{1, 2, 5, 2, 3}, target: 2},
			wantAns: []int{1, 2},
		},
		{
			args:    args{nums: []int{1, 2, 5, 2, 3}, target: 3},
			wantAns: []int{3},
		},
		{
			args:    args{nums: []int{1, 2, 5, 2, 3}, target: 5},
			wantAns: []int{4},
		},
		{
			args:    args{nums: []int{1, 2, 5, 2, 3}, target: 4},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("targetIndices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
