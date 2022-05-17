package p859

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: "ab", goal: "ab"},
			want: false,
		},
		{
			args: args{s: "ab", goal: "ba"},
			want: true,
		},
		{
			args: args{s: "aa", goal: "aa"},
			want: true,
		},
		{
			args: args{s: "aaaaaaabc", goal: "aaaaaaacb"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
