package p318

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			},
			want: 16,
		},
		{
			args: args{
				words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
			},
			want: 4,
		},
		{
			args: args{
				words: []string{"a", "aa", "aaa", "aaaa"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
