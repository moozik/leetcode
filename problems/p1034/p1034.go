package p1034

var dirs = []struct{ x, y int }{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	//边界点
	borders := []point{}
	originalColor := grid[row][col]
	//领地
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		//遍历四个方向
		isBorder := false
		for _, dir := range dirs {
			nx := x + dir.x
			ny := y + dir.y
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
			// if dirx == m || dirx < 0 || diry == n || diry < 0 {
			// 	isBorder = true
			// 	continue
			// }
			// fmt.Printf("vis:%+v,dirx:%+v,diry:%+v\n", vis, dirx, diry)
			// //领地
			// if vis[dirx][diry] {
			// 	continue
			// }
			// if grid[dirx][diry] != originalColor {
			// 	//边界
			// 	isBorder = true
			// 	dfs(dirx, diry)
			// 	continue
			// } else {
			// 	//新领地
			// 	vis[dirx][diry] = true
			// 	dfs(dirx, diry)
			// }
		}
		if isBorder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)
	// fmt.Printf("borders:%+v\n", borders)
	for _, point := range borders {
		grid[point.x][point.y] = color
	}
	return grid
}
