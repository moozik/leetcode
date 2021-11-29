package p786

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			args:    args{arr: []int{1, 2, 3, 5}, k: 3},
			wantAns: []int{2, 5},
		},
		{
			args:    args{arr: []int{1, 2, 3, 5, 7, 13, 71, 73}, k: 5},
			wantAns: []int{3, 73},
		},
		{
			args:    args{arr: []int{1, 2, 3, 5, 7, 13, 71, 73}, k: 6},
			wantAns: []int{3, 71},
		},
		{
			args:    args{arr: []int{1, 2, 3, 5, 7, 13, 71, 73}, k: 7},
			wantAns: []int{5, 73},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
