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

func fun(s string) (ret string) {
	for _, item := range s {
		if item >= '0' && item <= '9' {
			ret += string(item)
			continue
		}
		if item >= 'A' && item <= 'Z' {
			item = item + 32
			if item == 'z' {
				item = 'a'
			} else {
				item++
			}
			ret += string(item)
			continue
		}
		if item >= 'a' && item <= 'c' {
			ret += "2"
			continue
		}
		if item >= 'd' && item <= 'f' {
			ret += "3"
			continue
		}
		if item >= 'g' && item <= 'i' {
			ret += "4"
			continue
		}
		if item >= 'j' && item <= 'l' {
			ret += "5"
			continue
		}
		if item >= 'm' && item <= 'o' {
			ret += "6"
			continue
		}
		if item >= 'p' && item <= 's' {
			ret += "7"
			continue
		}
		if item >= 't' && item <= 'v' {
			ret += "8"
			continue
		}
		if item >= 'w' && item <= 'z' {
			ret += "9"
			continue
		}
	}
	return
}
