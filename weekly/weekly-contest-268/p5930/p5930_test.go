package p5930

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		colors []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			args:    args{colors: []int{1, 1, 1, 6, 1, 1, 1}},
			wantAns: 3,
		},
		{
			args:    args{colors: []int{1, 8, 3, 8, 3}},
			wantAns: 4,
		},
		{
			args:    args{colors: []int{0, 1}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxDistance(tt.args.colors); gotAns != tt.wantAns {
				t.Errorf("maxDistance() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
