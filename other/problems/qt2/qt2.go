package main

import (
	"fmt"
)

// 反转数字，不要使用string
func reverse(n int) (ret int) {
	for n != 0 {
		ret2 := n % 10
		ret = ret2 + ret*10
		n = int(n / 10)
	}
	fmt.Println(ret)
	return
}

func main() {
	fmt.Println(reverse(120) == 21)
	fmt.Println(reverse(123) == 321)
	fmt.Println(reverse(-123) == -321)
}
