package p794

import (
	"testing"
)

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{board: []string{
				"XOX", " X ", "   ",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTicTacToe(tt.args.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_win(t *testing.T) {
	type args struct {
		board []string
		b     byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{board: []string{"XXX", "   ", "OOO"}, b: 'O'},
			want: true,
		},
		{
			args: args{board: []string{"XXX", "   ", "OOO"}, b: 'X'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := win(tt.args.board, tt.args.b); got != tt.want {
				t.Errorf("win() = %v, want %v", got, tt.want)
			}
		})
	}
}
