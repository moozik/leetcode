package p517

import "fmt"

func findMinMoves(machines []int) (ans int) {
	n := len(machines)
	//衣服总数
	tot := 0
	for _, itemCount := range machines {
		tot += itemCount
	}
	//是否能均分
	if tot%n > 0 {
		return -1
	}
	//每个洗衣机里面的目标数量
	avg := tot / n
	//当前洗衣机需要移动的数量，>0为向右移动，<0为向左移动
	sum := 0
	for key, num := range machines {
		fmt.Printf("key:%v,num:%v,sum:%v,ans:%v\n", key, num, sum, ans)
		//计算当前与均值的差
		num -= avg
		//记录当前节点需要移动的数量
		sum += num
		//计算节点移动量 和 节点差值的max
		ans = max(ans, max(abs(sum), num))
		fmt.Printf("key:%v,num:%v,sum:%v,ans:%v\n\n", key, num, sum, ans)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

