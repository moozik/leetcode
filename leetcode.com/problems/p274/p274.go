package p274

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	maxH := 0
	for i, item := range citations {
		if item == 0 {
			continue
		}
		if item < n-i {
			continue
		}
		m := min(item, n-i)
		if m < maxH {
			break
		}
		maxH = max(m, maxH)
	}
	return maxH
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
