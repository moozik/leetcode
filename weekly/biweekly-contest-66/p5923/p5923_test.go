package p5923

import "testing"

func Test_minimumBuckets(t *testing.T) {
	type args struct {
		street string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			args:    args{street: ".H.H.H.H."},
			wantAns: 2,
		},
		{
			args:    args{street: "...H.H.H..H.H..H.H..."},
			wantAns: 4,
		},
		{
			args:    args{street: ".H.H.H.H.H."},
			wantAns: 3,
		},
		{
			args:    args{street: "H..H"},
			wantAns: 2,
		},
		{
			args:    args{street: ".HHH."},
			wantAns: -1,
		},
		{
			args:    args{street: ".H.H.H.H.H."},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumBuckets(tt.args.street); gotAns != tt.wantAns {
				t.Errorf("minimumBuckets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
