package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		printArr(fun(input.Text()))
	}
}

func printArr(data []string) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func fun(data string) (ret []string) {
	if len(data) <= 8 {
		if len(data) < 8 {
			data = data + strings.Repeat("0", 8-len(data))
		}
		ret = append(ret, data)
		return
	}

	math.Floor(float64(len(data)) / 8)

	for i := 0; i < len(data); i += 8 {
		var item string
		if len(data)-i < 8 {
			item = data[i:]
			item = item + strings.Repeat("0", 8-(len(data)-i))
		} else {
			item = data[i : i+8]
		}
		ret = append(ret, item)
	}
	// fmt.Println(ret)
	return
}
