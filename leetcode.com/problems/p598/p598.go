package p598

func maxCount(m int, n int, ops [][]int) int {
	maxM := m
	maxN := n
	for _, item := range ops {
		maxM = min(maxM, item[0])
		maxN = min(maxN, item[1])
	}
	return maxM * maxN
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
