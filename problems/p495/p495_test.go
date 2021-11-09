package p495

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	type args struct {
		timeSeries []int
		duration   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				timeSeries: []int{1, 4},
				duration:   2,
			},
			want: 4,
		},
		{
			args: args{
				timeSeries: []int{1, 2},
				duration:   2,
			},
			want: 3,
		},
		{
			args: args{
				timeSeries: []int{1, 2, 3, 4, 5},
				duration:   5,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
