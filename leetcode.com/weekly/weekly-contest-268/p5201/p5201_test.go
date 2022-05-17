package p5201

import "testing"

func Test_wateringPlants(t *testing.T) {
	type args struct {
		plants   []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{plants: []int{2, 2, 3, 3}, capacity: 5},
			want: 14,
		},
		{
			args: args{plants: []int{1, 1, 1, 4, 2, 3}, capacity: 4},
			want: 30,
		},
		{
			args: args{plants: []int{7, 7, 7, 7, 7, 7, 7}, capacity: 8},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStepCount := wateringPlants(tt.args.plants, tt.args.capacity); gotStepCount != tt.want {
				t.Errorf("wateringPlants() = %v, want %v", gotStepCount, tt.want)
			}
		})
	}
}
