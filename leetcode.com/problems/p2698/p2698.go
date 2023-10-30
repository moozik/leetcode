package p2698

import "strconv"

func punishmentNumber(n int) int {
	ret := 0
	for i := 9; i <= n; i++ {
		if dp(strconv.Itoa(i*i), i) == i {
			ret += i * i
		}
	}
	return ret + 1
}

func dp(s string, target int) int {
	a, _ := strconv.Atoi(s)
	if a < 10 {
		return a
	}
	if a == target {
		return target
	}
	for k := range s {
		val, _ := strconv.Atoi(s[:k])
		if dp(s[k:], target-val) == target-val {
			return target
		}
	}
	return 0
}
