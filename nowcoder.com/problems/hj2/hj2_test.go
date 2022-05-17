package main

import "testing"

func Test_fun(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			args: args{
				str1: "34a5435 45a34 54a35 43a5 a34 534 534",
				str2: "a",
			},
			wantRet: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := fun(tt.args.str1, tt.args.str2); gotRet != tt.wantRet {
				t.Errorf("fun() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
