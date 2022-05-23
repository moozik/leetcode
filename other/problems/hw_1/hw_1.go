package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

func main() {
	index := 0
	command := ""
	for {
		_, err := fmt.Scan(&index)
		if err == io.EOF {
			break
		}
		fmt.Scan(&command)
		fmt.Println(fun(index, command))
	}
}

func fun(index int, command string) (ret string) {
	// fmt.Println(index, command)
	command = strings.Trim(command, "_")
	r, _ := regexp.Compile(`("[A-Za-z0-9_]*?"|[A-Za-z0-9]*)`)
	res := r.FindAllStringSubmatch(command, -1)
	if index >= len(res) || index < 0 {
		return "ERROR"
	}
	// fmt.Println(res)
	k := 0
	for _, item := range res {
		if item[0] == "" {
			continue
		}
		if k != index {
			ret = ret + "_" + item[0]
			k++
			continue
		}
		ret = ret + "_******"
		k++
	}
	return strings.Trim(ret, "_")
}
