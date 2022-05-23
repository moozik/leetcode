package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	nums := make([]int, b)
	for i := 0; i < 10; i++ {
		item := 0
		fmt.Scan(&item)
		nums[add(item)%b]++
	}
	maxNum := -1
	for k, item := range nums {
		if k < a && item > maxNum {
			maxNum = item
		}
	}
	fmt.Println(maxNum)
	// fmt.Println(a, b, nums)
}

func add(in int) (ret int) {
	for i := 0; i < 4; i++ {
		ret = ret + (in % 255)
		in = in >> 32
	}
	return
}
