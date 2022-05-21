package main

import "testing"

func Test_fun(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			args: args{
				"5VQ6HnjWx8GCg6",
			},
			//5wr6i6x98hd46
			//5wr6i65x98hd46
			wantRet: "5wr6i65x98hd46",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := fun(tt.args.s); gotRet != tt.wantRet {
				t.Errorf("fun() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
