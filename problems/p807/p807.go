package p807

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	maxHeight := make([][]int, 2)
	maxHeight[0], maxHeight[1] = make([]int, len(grid[0])), make([]int, len(grid))

	for y, rows := range grid {
		maxTmp := 0
		for x, height := range rows {
			maxTmp = max(maxTmp, height)
			maxHeight[0][x] = max(maxHeight[0][x], height)
		}
		maxHeight[1][y] = maxTmp
	}

	for y, rows := range grid {
		for x, height := range rows {
			minHeight := min(maxHeight[0][x], maxHeight[1][y])
			ans += minHeight - height
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
