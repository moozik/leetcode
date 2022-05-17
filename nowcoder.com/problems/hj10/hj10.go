package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(fun(str))
}

func fun(str string) (ret int) {
	// fmt.Println(str)
	save := make(map[rune]bool)
	for _, item := range str {
		if item == 10 || item == 13 {
			continue
		}
		// fmt.Println(item)
		if save[item] {
			continue
		}
		save[item] = true
		ret++
	}
	return
}
