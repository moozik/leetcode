package p1816

import "testing"

func Test_truncateSentence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "Hello how are you Contestant", k: 1},
			want: "Hello",
		},
		{
			args: args{s: "Hello how are you Contestant", k: 4},
			want: "Hello how are you",
		},
		{
			args: args{s: "What is the solution to this problem", k: 4},
			want: "What is the solution",
		},
		{
			args: args{s: "chopper is not a tanuki", k: 5},
			want: "chopper is not a tanuki",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateSentence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
