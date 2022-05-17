package p5924

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		startPos []int
		homePos  []int
		rowCosts []int
		colCosts []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			args: args{
				startPos: []int{1, 0},
				homePos:  []int{2, 3},
				rowCosts: []int{5, 4, 3},
				colCosts: []int{8, 2, 6, 7},
			},
			wantAns: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minCost(tt.args.startPos, tt.args.homePos, tt.args.rowCosts, tt.args.colCosts); gotAns != tt.wantAns {
				t.Errorf("minCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
