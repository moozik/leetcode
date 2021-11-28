package p5941

import (
	"reflect"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
	type args struct {
		n           int
		meetings    [][]int
		firstPerson int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			args: args{n: 6, meetings: [][]int{
				{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
			}, firstPerson: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAllPeople(tt.args.n, tt.args.meetings, tt.args.firstPerson); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAllPeople() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
