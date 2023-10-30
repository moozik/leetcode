package p2520

import "testing"

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{num: 7},
			want: 1,
		},
		{
			args: args{num: 121},
			want: 2,
		},
		{
			args: args{num: 1248},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.num); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
