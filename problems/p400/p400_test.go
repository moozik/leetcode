package p400

import (
	"testing"
)

func Test_findNthDigitAddCase(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 5},
			want: 5,
		},
		{
			args: args{n: 11},
			want: 0,
		},
		{
			args: args{n: 100},
			want: 5,
		},
		{
			args: args{n: 200},
			want: 0,
		},
		{
			args: args{n: 1102},
			want: 4,
		},
		{
			args: args{n: 1103},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigitAddCase(tt.args.n); got != tt.want {
				t.Errorf("findNthDigitAddCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 5},
			want: 5,
		},
		{
			args: args{n: 11},
			want: 0,
		},
		{
			args: args{n: 100},
			want: 5,
		},
		{
			args: args{n: 200},
			want: 0,
		},
		{
			args: args{n: 1102},
			want: 4,
		},
		{
			args: args{n: 1103},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
