package p594

func findLHS(nums []int) (ans int) {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}
	for num, count := range numCountMap {
		if count2, ok := numCountMap[num+1]; ok && count+count2 > ans {
			ans = count + count2
		}
	}
	return ans
}
