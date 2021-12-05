package p372

const mod = 1337

func superPow(a int, b []int) (ans int) {
	a = a % 1337
	ans = 1
	//低位到高位
	for i := len(b) - 1; i > -1; i-- {
		//求解
		ans = ans * quickPow(a, b[i]) % mod
		a = quickPow(a, 10) % mod
	}
	return
}

func quickPow(x int, n int) (ans int) {
	ans = 1
	x_contribute := x
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			ans = ans * x_contribute % mod
		}
		x_contribute = x_contribute * x_contribute % mod
	}
	return
}
