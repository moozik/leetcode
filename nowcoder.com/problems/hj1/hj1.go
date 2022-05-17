package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(fun(msg))
}

func fun(msg string) int {
	msg = strings.TrimSpace(msg)

	list1 := strings.Fields(msg)
	word := list1[len(list1)-1]
	return len(word)
}
