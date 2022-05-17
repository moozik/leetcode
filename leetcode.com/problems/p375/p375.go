package p375

import (
	// "fmt"
	"math"
)

//https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/

func getMoneyAmount(n int) (ans int) {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for leftPos := n - 1; leftPos > 0; leftPos-- {
		for rightPos := leftPos + 1; rightPos <= n; rightPos++ {
			ans = math.MaxInt32
			for k := leftPos; k < rightPos; k++ {
				ret := k + max(f[leftPos][k-1], f[k+1][rightPos])
				if ret < ans {
					ans = ret
				}
			}
			f[leftPos][rightPos] = ans
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
