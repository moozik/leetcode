package p42_2

//官方解法

func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return 0
	}
	//获取左向max
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for index := 1; index < n; index++ {
		leftMax[index] = max(leftMax[index-1], height[index])
	}
	//获取右向max
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for index := n - 2; index >= 0; index-- {
		rightMax[index] = max(rightMax[index+1], height[index])
	}

	for key, h := range height {
		ans += min(leftMax[key], rightMax[key]) - h
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
