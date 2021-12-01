package p1446

func maxPower(s string) (ans int) {
	lastStr := s[0]
	count := 1
	n := len(s)
	ans = 1
	for i := 1; i < n; i++ {
		if s[i] != lastStr {
			count = 1
			lastStr = s[i]
			continue
		}
		count++
		if count > ans {
			ans = count
		}
	}
	return
}
