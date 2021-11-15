//https://leetcode-cn.com/problems/bulb-switcher/
package p319

import (
	"fmt"
	"math"
)

//超时解法
func bulbSwitch_timeout(n int) (ans int) {
	fmt.Printf("start,n:%+v\n", n)
	if n <= 3 {
		return 1
	}
	ans = 1
	for i := 4; i <= n; i++ {
		flagInt := 0
		for k := 2; k <= i>>1; k++ {
			fmt.Printf("i:%+v,k:%+v\n", i, k)
			if i%k == 0 {
				flagInt++
			}
		}
		if flagInt%2 == 1 {
			fmt.Printf("pos,i:%+v,brint\n", i)
			ans++
		}
	}
	return
}

//超时解法2 写了俩都超时心态炸了，题解就一行，心态更炸了
func bulbSwitch_timeout2(n int) (ans int) {
	fmt.Printf("start,n:%+v\n", n)
	if n <= 3 {
		return 1
	}
	ans = 1
	//奇数代表亮灯
	lightStage := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for k := 1; k*i <= n; k++ {
			fmt.Printf("k:%+v,i:%+v,k*i:%+v\n", k, i, k*i)
			lightStage[k*i]++
		}
		if lightStage[i]%2 == 0 {
			ans++
		}
	}
	return
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
