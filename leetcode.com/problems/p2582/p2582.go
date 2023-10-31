package p2582

func passThePillow(n int, time int) int {
	time++
	time2 := time % (2*n - 2)
	if time2 == 0 {
		return 2
	}
	if time2 < n {
		return time2
	}
	return n*2 - time2
}
