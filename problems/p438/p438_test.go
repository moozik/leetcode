package p438

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			args:    args{s: "cbaebabacd", p: "abc"},
			wantAns: []int{0, 6},
		},
		{
			args:    args{s: "abab", p: "ab"},
			wantAns: []int{0, 1, 2},
		},
		{
			args:    args{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", p: "aaa"},
			wantAns: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		},
		{
			args:    args{s: "aaaaaaaaaa", p: "aaaaaaaaaaaaa"},
			wantAns: []int{},
		},
		{
			args:    args{s: "aaaaaaaaaaaaa", p: "aaaaaaaaaaaaa"},
			wantAns: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAnagrams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
