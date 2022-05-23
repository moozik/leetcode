package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			args:    args{256},
			wantRet: 1,
		},
		{
			args:    args{257},
			wantRet: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := add(tt.args.in); gotRet != tt.wantRet {
				t.Errorf("add() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
