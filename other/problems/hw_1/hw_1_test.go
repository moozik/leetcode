package main

import "testing"

func Test_fun(t *testing.T) {
	type args struct {
		index   int
		command string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			args: args{
				index:   1,
				command: `password_"_a12345678"_timeout_100`,
			},
			wantRet: "password_******_timeout_100",
		},
		{
			args: args{
				index:   1,
				command: "password__a12345678_timeout_100",
			},
			wantRet: "password_******_timeout_100",
		},
		{
			args: args{
				index:   2,
				command: `aaa_password_"a12_45678"_timeout__100_""_`,
			},
			wantRet: `aaa_password_******_timeout_100_""`,
		},
		{
			args: args{
				index:   1,
				command: "___password__a12345678_timeout_100___",
			},
			wantRet: "password_******_timeout_100",
		},
		{
			args: args{
				index:   0,
				command: `password`,
			},
			wantRet: "******",
		},
		{
			args: args{
				index:   2,
				command: `password`,
			},
			wantRet: "ERROR",
		},
		{
			args: args{
				index:   0,
				command: `__"1"_""_password_"______"__`,
			},
			wantRet: `******_""_password_"______"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := fun(tt.args.index, tt.args.command); gotRet != tt.wantRet {
				t.Errorf("fun() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
