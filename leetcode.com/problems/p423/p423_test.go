package p423

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "owoztneoer"},
			want: "012",
		},
		{
			args: args{s: "fviefuro"},
			want: "45",
		},
		{
			args: args{s: "nnei"},
			want: "9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
