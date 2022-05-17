package p5940

func minimumDeletions(nums []int) int {
	n := len(nums)
	minPos, maxPos := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] > nums[maxPos] {
			maxPos = i
		}
		if nums[i] < nums[minPos] {
			minPos = i
		}
	}
	// fmt.Printf("minPos:%+v, maxPos:%+v\n", minPos, maxPos)

	if minPos > maxPos {
		minPos, maxPos = maxPos, minPos
	}
	// fmt.Printf("minPos:%+v, maxPos:%+v\n", minPos, maxPos)
	//1 5
	//2 + (8 - 5）
	//5
	// fmt.Printf("minPos+1：%+v\n", minPos+1)
	// fmt.Printf("(n-maxPos)：%+v\n", (n - maxPos))
	// fmt.Printf("maxPos+1：%+v\n", maxPos+1)
	// fmt.Printf("n-minPos：%+v\n", n-minPos)
	return min(minPos+1+(n-maxPos), min(maxPos+1, n-minPos))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
