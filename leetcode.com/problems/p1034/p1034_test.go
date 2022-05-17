package p1034

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		row   int
		col   int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				grid:  [][]int{{1, 1}, {1, 2}},
				row:   0,
				col:   0,
				color: 3,
			},
			want: [][]int{{3, 3}, {3, 2}},
		},
		{
			args: args{
				grid:  [][]int{{1, 2, 2}, {2, 3, 2}},
				row:   0,
				col:   1,
				color: 3,
			},
			want: [][]int{{1, 3, 3}, {2, 3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorBorder(tt.args.grid, tt.args.row, tt.args.col, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
