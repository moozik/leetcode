package p2605

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums1: []int{4, 1, 3},
				nums2: []int{5, 7},
			},
			want: 15,
		},
		{
			args: args{
				nums1: []int{3, 5, 2, 6},
				nums2: []int{3, 1, 7},
			},
			want: 3,
		},
		{
			args: args{
				nums1: []int{6, 4, 3, 7, 2, 1, 8, 5},
				nums2: []int{6, 8, 5, 3, 1, 7, 4, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
