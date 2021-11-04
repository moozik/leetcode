package p42

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{height: []int{0,1,0,2,1,0,1,3,2,1,2,1}},
			wantRet: 6,
		},
		{
			name: "2",
			args: args{height: []int{4,2,0,3,2,5}},
			wantRet: 9,
		},
		{
			name: "3",
			args: args{height: []int{3,1,4,7,5,0,1,5,1,5,9,6,8,9,3,3,1}},
			wantRet: 31,
		},
		{
			name: "4",
			args: args{height: []int{0,4,5,7,6,7,5,7,5,0,7,7,0,9,9,3,9,9,3}},
			wantRet: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := trap(tt.args.height); gotRet != tt.wantRet {
				t.Errorf("trap() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}