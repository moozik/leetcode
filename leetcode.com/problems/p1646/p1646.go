package p1646

func getMaximumGenerated(n int) int {
	if 1 == n || 2 == n {
		return 1
	}
	nums := make([]int, n+1)
	nums[1] = 1
	max := 0
	for i := 1; i*2 <= n; i++ {
		if i*2 <= n {
			nums[i*2] = nums[i]
		}
		if i*2+1 <= n {
			nums[i*2+1] = nums[i] + nums[i+1]
			if nums[i*2+1] > max {
				max = nums[i*2+1]
			}
		}
	}
	// fmt.Println(nums)
	return max
}
