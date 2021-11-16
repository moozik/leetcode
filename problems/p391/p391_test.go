//https://leetcode-cn.com/problems/perfect-rectangle/
package p391

import "testing"

func Test_isRectangleCover(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	args: args{rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}},
		// 	want: true,
		// },
		// {
		// 	args: args{rectangles: [][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}},
		// 	want: false,
		// },
		// {
		// 	args: args{rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}}},
		// 	want: false,
		// },
		// {
		// 	args: args{rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}},
		// 	want: false,
		// },
		// {
		// 	//内部覆盖
		// 	args: args{rectangles: [][]int{{0, 0, 1, 1}, {0, 1, 3, 2}, {1, 0, 2, 2}}},
		// 	want: false,
		// },
		{
			//覆盖
			args: args{rectangles: [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}, {1, 1, 2, 2}, {1, 1, 2, 2}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleCover(tt.args.rectangles); got != tt.want {
				t.Errorf("isRectangleCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
