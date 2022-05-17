package p400

import (
	"fmt"
	"math"
	"strconv"
)

//https://leetcode-cn.com/problems/nth-digit/
func findNthDigit(n int) int {
	var num float64
	if n < 10 {
		return n
	}
	for i := 1; i <= 31; i++ {
		num = math.Pow(10, float64(i))
		if num <= float64(n/i) {
			// fmt.Printf("pos1_num:%+v,n:%+v,i:%+v\n", num, n, i)
			//n往上加
			n = n + int(num)
			continue
		}
		// fmt.Printf("pos2_num:%+v,n:%+v,i:%+v\n", num, n, i)
		//算答案
		targetNum := strconv.Itoa(n / i)
		targetPos := n % i
		// fmt.Printf("targetNum:%+v,targetPos:%+v\n", targetNum, targetPos)
		ans, _ := strconv.Atoi(string(targetNum[targetPos]))
		return ans
	}
	return 0
}

func findNthDigitAddCase(n int) int {
	numLine := ""
	for i := 1; i > 0; i++ {
		if len(numLine) >= n {
			num, _ := strconv.Atoi(string(numLine[n-1]))
			fmt.Printf("numLine:%+v\n", numLine)
			fmt.Printf("n:%+v,ans:%+v\n", n, num)
			return num
		}
		numLine = numLine + strconv.Itoa(i)
	}
	return 0
}
