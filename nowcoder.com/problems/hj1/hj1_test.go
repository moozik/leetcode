package main

import "testing"

func Test_fun(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				msg: "hello nowcoder",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fun(tt.args.msg); got != tt.want {
				t.Errorf("fun() = %v, want %v", got, tt.want)
			}
		})
	}
}
