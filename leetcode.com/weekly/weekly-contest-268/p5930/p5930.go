package p5930

func maxDistance(colors []int) (ans int) {
	earlyMap := map[int]int{}
	lastMap := map[int]int{}
	for index, col := range colors {
		if _, ok := earlyMap[col]; !ok {
			earlyMap[col] = index
		}
		if item, ok := lastMap[col]; !ok || item < index {
			lastMap[col] = index
		}
	}
	// fmt.Printf("earlyMap:%+v\n", earlyMap)
	// fmt.Printf("lastMap:%+v\n", lastMap)
	for col, pos := range earlyMap {
		for col2, pos2 := range lastMap {
			if col == col2 {
				continue
			}
			// fmt.Printf("abs(pos-pos2):%+v\n", abs(pos-pos2))
			ans = max(ans, abs(pos-pos2))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
