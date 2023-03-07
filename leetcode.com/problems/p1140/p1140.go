package p1140

import (
	"fmt"
)

func stoneGameII(piles []int) int {
	ai, _ := dp(true, 0, 0, 1, piles)
	return ai
}

func dp(aliceRound bool, ai, bo, m int, piles []int) (int, int) {
	fmt.Printf("[dp入口]round:%+v, ai:%+v, bo:%+v, m:%+v, piles:%+v\n", aliceRound, ai, bo, m, piles)
	if len(piles) == 0 {
		return ai, bo
	}
	n := len(piles)
	if m*2 >= n {
		//剩下的可以全拿走
		if aliceRound {
			fmt.Printf("[dp返回2]m:%+v,ai:%+v,bo%+v\n\n", m, ai+sum(piles), bo)
			return ai + sum(piles), bo
		} else {
			fmt.Printf("[dp返回2]m:%+v,ai:%+v,bo%+v\n\n", m, ai, bo+sum(piles))
			return ai, bo + sum(piles)
		}
	}
	aiBest := 0
	boBest := 0
	for i := 1; i <= m*2 && i <= n; i++ {
		if aliceRound {
			aiTmp, _ := dp(!aliceRound, ai+sum(piles[:i]), bo, max(m, i), piles[i:])
			aiBest = max(aiTmp, aiBest)
			fmt.Printf("ai Round,m:%+v,i:%+v, aiTmp:%+v\n", m, i, aiTmp)
		} else {
			_, boTmp := dp(!aliceRound, ai, bo+sum(piles[:i]), max(m, i), piles[i:])
			boBest = max(boTmp, boBest)
			fmt.Printf("bo Round,m:%+v,i:%+v, aiTmp:%+v\n", m, i, boTmp)
		}
	}
	fmt.Printf("[dp返回]m:%+v,ai:%+v,bo%+v\n\n", m, aiBest+ai, boBest+bo)
	return aiBest + ai, boBest + bo
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
