package p2582

import "testing"

func Test_passThePillow(t *testing.T) {
	type args struct {
		n    int
		time int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:    18,
				time: 38,
			},
			want: 5,
		},
		{
			args: args{
				n:    4,
				time: 5,
			},
			want: 2,
		},
		{
			args: args{
				n:    3,
				time: 2,
			},
			want: 3,
		},
		{
			args: args{
				n:    20,
				time: 5,
			},
			want: 6,
		},
		{
			args: args{
				n:    20,
				time: 22,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := passThePillow(tt.args.n, tt.args.time); got != tt.want {
				t.Errorf("passThePillow() = %v, want %v", got, tt.want)
			}
		})
	}
}
