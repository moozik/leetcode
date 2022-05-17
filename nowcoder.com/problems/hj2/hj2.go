package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	//读取一行
	input.Scan()
	str1 := input.Text() // 获取第一行的数据

	//读取第二行
	input.Scan()
	str2 := input.Text()

	fmt.Println(fun(str1, str2))
}

func fun(str1 string, str2 string) (ret int) {
	str2 = strings.ToLower(str2)
	for _, item := range strings.ToLower(str1) {
		if string(item) == str2 {
			ret++
		}
	}
	return
}
