package p458

import "testing"

func Test_poorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//一个单位时间
		{
			args: args{buckets: 1000, minutesToDie: 15, minutesToTest: 15},
			want: 10,
		},
		{
			args: args{buckets: 200, minutesToDie: 15, minutesToTest: 15},
			want: 8,
		},
		//多个单位时间
		{
			args: args{buckets: 1000, minutesToDie: 15, minutesToTest: 60},
			want: 5,
		},
		{
			args: args{buckets: 4, minutesToDie: 15, minutesToTest: 15},
			want: 2,
		},
		{
			args: args{buckets: 4, minutesToDie: 15, minutesToTest: 30},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest); got != tt.want {
				t.Errorf("poorPigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
