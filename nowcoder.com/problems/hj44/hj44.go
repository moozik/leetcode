package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sudoku [][]int = make([][]int, 9)
var sign bool = false

func main() {
	// fmt.Println(int(80 / 9))
	sudoku = inputIntArr()
	dfs(0)
	printArr(sudoku)
}

func printArr(out [][]int) {
	for _, item := range out {
		for _, item2 := range item {
			fmt.Printf("%d ", item2)
		}
		fmt.Println()
	}
}

func dfs(index int) {
	row := int(index / 9)
	col := index % 9
	if index >= 81 {
		sign = true
		//返回点
		return
	}
	if sudoku[row][col] != 0 {
		dfs(index + 1)
		return
	}
	for i := 1; i <= 9; i++ {
		if !isVaild(i, row, col) {
			continue
		}
		sudoku[row][col] = i
		dfs(index + 1)
		if sign {
			return
		}
	}
	sudoku[row][col] = 0
}

func isVaild(num, x, y int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[x][i] == num || sudoku[i][y] == num {
			return false
		}
		yy := int(x/3)*3 + int(i/3)
		xx := int(y/3)*3 + i%3
		// fmt.Printf("yy:%d, xx:%d\n", yy, xx)
		if sudoku[yy][xx] == num {
			return false
		}
	}
	return true
}

func inputIntArr() (ret [][]int) {
	input := bufio.NewScanner(os.Stdin)
	i := 0
	ret = make([][]int, 9)
	for input.Scan() {
		in := input.Text()
		if in == "" {
			return
		}
		ret[i] = make([]int, 0)
		arr := strings.Split(in, " ")
		for k, item := range arr {
			if k >= 9 {
				continue
			}
			num, _ := strconv.Atoi(item)
			ret[i] = append(ret[i], num)
		}
		i++
	}
	return
}
