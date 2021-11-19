package p397

//贪心算法最优解 参考了题解
func integerReplacement(n int) (ans int) {
	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
			ans++
			continue
		}
		if n%4 == 1 {
			n = n - 1
			ans++
			continue
		}
		if n == 3 {
			ans += 2
			break
		}
		n = n + 1
		ans++
	}
	return ans
}
