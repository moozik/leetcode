//快速幂
//https://leetcode-cn.com/problems/powx-n/
package p50

func myPow(x float64, n int) (ans float64) {
	if n < 0 {
		return 1.0 / quickMul(x, -n)
	}
	return quickMul(x, n)
}

func quickMul(x float64, n int) (ans float64) {
	ans = 1
	x_contribute := x
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			ans *= x_contribute
			// fmt.Printf("ans:%+v\n", ans)
		}
		x_contribute *= x_contribute
	}
	return
}
