package p520

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantAns bool
	}{
		{
			args:    args{word: "USA"},
			wantAns: true,
		},
		{
			args:    args{word: "LEETCODE"},
			wantAns: true,
		},
		{
			args:    args{word: "wORD"},
			wantAns: false,
		},
		{
			args:    args{word: "Word"},
			wantAns: true,
		},
		{
			args:    args{word: "AbA"},
			wantAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := detectCapitalUse(tt.args.word); gotAns != tt.wantAns {
				t.Errorf("detectCapitalUse() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
