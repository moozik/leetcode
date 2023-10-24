package p1103

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies    int
		num_people int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				candies:    60,
				num_people: 4,
			},
			want: []int{15, 18, 15, 12},
		},
		{
			args: args{
				candies:    7,
				num_people: 4,
			},
			want: []int{1, 2, 3, 1},
		},
		{
			args: args{
				candies:    10,
				num_people: 3,
			},
			want: []int{5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
