package p5939

//https://leetcode-cn.com/problems/k-radius-subarray-averages/submissions/
func getAverages(nums []int, k int) (ans []int) {
	n := len(nums)
	if n < k*2+1 {
		for i := 0; i < n; i++ {
			ans = append(ans, -1)
		}
		return ans
	}
	sumCount := 0

	if k != 0 {
		for o := 0; o < k*2; o++ {
			sumCount += nums[o]
		}
	}
	for i := 0; i < n; i++ {
		if k != 0 && (i < k || i > n-k-1) {
			ans = append(ans, -1)
			continue
		}
		if k != 0 {
			// fmt.Printf("sumCount:%+v,i-k-1:%+v,i+k:%+v\n", sumCount, i-k-1, i+k)

			if i-k-1 >= 0 {
				sumCount -= nums[i-k-1]
			}
			if i+k < n {
				sumCount += nums[i+k]
			}
			// fmt.Printf("sumCount:%+v\n", sumCount)
			ans = append(ans, sumCount/(k*2+1))
		} else {
			ans = append(ans, nums[i])
		}
	}
	return ans
}
