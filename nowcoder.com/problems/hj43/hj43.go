package main

import "fmt"

/*
输入
5 5
0 1 0 0 0
0 1 1 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0
*/

var result [][]int
var nums [][]int
var m, n int

func main() {
	result = make([][]int, 0)
	fmt.Scan(&m, &n)

	nums = make([][]int, m)

	for i := 0; i < m; i++ {
		nums[i] = make([]int, n)
		for o := 0; o < n; o++ {
			fmt.Scan(&nums[i][o])
		}
	}
	// fmt.Println(nums)
	dfs(0, 0)
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Printf("(%d,%d)\n", result[i][0], result[i][1])
	}
}

func dfs(x, y int) bool {
	// fmt.Println(x, y)
	//非法
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	//墙
	if nums[x][y] == 1 {
		return false
	}
	//结束
	if nums[x][y] == 0 && x == m-1 && y == n-1 {
		result = append(result, []int{x, y})
		return true
	}
	//把来的时候的路堵上
	nums[x][y] = 1
	//上下左右
	if x < m-1 && dfs(x+1, y) {
		result = append(result, []int{x, y})
		return true
	}
	if y < n-1 && dfs(x, y+1) {
		result = append(result, []int{x, y})
		return true
	}
	if x > 0 && dfs(x-1, y) {
		result = append(result, []int{x, y})
		return true
	}
	if y > 0 && dfs(x, y-1) {
		result = append(result, []int{x, y})
		return true
	}
	return false
}
