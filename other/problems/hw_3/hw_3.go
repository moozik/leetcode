package main

import (
	"fmt"
	"strings"
)

var result map[string]bool
var text []string
var targetLen int

func main() {
	index := 0
	// for {
	result = make(map[string]bool)
	fmt.Scan(&index)
	text = make([]string, index)
	for i := 0; i < index; i++ {
		fmt.Scan(&text[i])
	}
	textJoin := strings.Join(text, "")
	targetLen = len(textJoin)
	fun()
	// fmt.Println(result)
	resString := []string{}
	for k, _ := range result {
		resString = append(resString, k)
		// fmt.Println(k)
	}
	// fmt.Println(resString)
	for i := 0; i < len(resString)-1; i++ {
		for o := i + 1; o < len(resString); o++ {
			if resString[o] < resString[i] {
				resString[o], resString[i] = resString[i], resString[o]
				continue
			}
		}
	}
	// fmt.Println(resString)
	for _, item := range resString {
		fmt.Println(item)
	}
	// }
}

func fun() {
	// fmt.Println(text)
	dp("", []int{})
}

func dp(res string, usedIndex []int) {
	// fmt.Println(res, usedIndex, result)
	if len(res) == targetLen {
		result[res] = true
		return
	}
	usedLen := len(usedIndex)
	usedIndex = append(usedIndex, -1)
	for k, item := range text {
		// fmt.Println(k, item)
		if isHave(k, &usedIndex) {
			continue
		}
		usedIndex[usedLen] = k
		dp(res+item, usedIndex)
	}
}

func isHave(target int, nums *[]int) bool {
	for _, item := range *nums {
		if target == item {
			return true
		}
	}
	return false
}
