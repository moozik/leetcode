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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
