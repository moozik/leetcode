package p629

import "testing"

func Test_kInversePairs(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 1000,
				k: 1000,
			},
			want: 663677020,
		},
		//{
		//	args: args{
		//		n: 25,
		//		k: 24,
		//	},
		//	want: 197331137,
		//},
		//{
		//	args: args{
		//		n: 89,
		//		k: 1,
		//	},
		//	want: 88,
		//},
		//{
		//	args: args{
		//		n: 5,
		//		k: 1,
		//	},
		//	want: 4,
		//},
		//{
		//	args: args{
		//		n: 4,
		//		k: 2,
		//	},
		//	want: 5,
		//},
		//{
		//	args: args{
		//		n: 5,
		//		k: 2,
		//	},
		//	want: 9,
		//},
		//{
		//	args: args{
		//		n: 5,
		//		k: 7,
		//	},
		//	want: 15,
		//},
		//{
		//	args: args{
		//		n: 5,
		//		k: 9,
		//	},
		//	want: 4,
		//},
		//{
		//	args: args{
		//		n: 10,
		//		k: 2,
		//	},
		//	want: 44,
		//},
		//{
		//	args: args{
		//		n: 10,
		//		k: 10,
		//	},
		//	want: 21670,
		//},
		//{
		//	args: args{
		//		n: 3,
		//		k: 0,
		//	},
		//	want: 1,
		//},
		//{
		//	args: args{
		//		n: 3,
		//		k: 1,
		//	},
		//	want: 2,
		//},
		//{
		//	args: args{
		//		n: 25,
		//		k: 3,
		//	},
		//	want: 2575,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kInversePairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
