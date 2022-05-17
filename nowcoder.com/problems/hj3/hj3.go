package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	printArr(fun(inputIntArr()))
}

func inputIntArr() []int {
	input := bufio.NewScanner(os.Stdin)
	data := make([]int, 0)
	for input.Scan() {
		in := input.Text()
		if in == "" {
			return data
		}
		intData, _ := strconv.Atoi(in)
		data = append(data, intData)
	}
	return data
}

func printArr(data []int) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

type Arr []int

func (t Arr) Len() int {
	return len(t)
}
func (t Arr) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func fun(arr []int) (ret []int) {
	var arr2 Arr = append([]int(nil), arr[1:]...)
	sort.Sort(arr2)
	for k, item := range arr2 {
		if k+1 == len(arr2) || item != arr2[k+1] {
			ret = append(ret, item)
		}
	}
	return ret
}
