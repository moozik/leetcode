package p5938

func targetIndices(nums []int, target int) (ans []int) {
	minCount := 0
	total := 0
	for _, item := range nums {
		if item < target {
			minCount++
		}
		if item == target {
			total++
		}
	}
	if total == 0 {
		return ans
	}
	ans = []int{}
	for i := 0; i < total; i++ {
		ans = append(ans, minCount+i)
	}
	return ans
}
