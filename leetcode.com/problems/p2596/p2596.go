package p2596

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	yLen := len(grid)
	if yLen == 3 {
		return false
	}
	xLen := len(grid[0])
	save := make([][]int, yLen*xLen)
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			save[grid[y][x]] = []int{y, x}
		}
	}
	for i, item := range save {
		if i == 0 {
			continue
		}
		if Abs(save[i-1][0]-item[0])*Abs(save[i-1][1]-item[1]) != 2 {
			return false
		}
	}
	return true
}

func Abs[T int8 | int32 | int64 | int | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
