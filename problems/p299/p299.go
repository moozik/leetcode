package p299

import (
	"fmt"
)

//https://leetcode-cn.com/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	n := len(secret)
	bullsCount := 0
	cowsCount := 0
	secretMap := [10]int{}
	guessMap := [10]int{}
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			bullsCount++
		} else {
			secretMap[secret[i]-'0']++
			guessMap[guess[i]-'0']++
		}
	}
	for index, item := range guessMap {
		cowsCount += min(item, secretMap[index])
	}
	return fmt.Sprintf("%dA%dB", bullsCount, cowsCount)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
