package p1

func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 < k2 {
				if v1+v2 == target {
					return []int{k1, k2}
				}
			}
		}
	}
	return []int{}
}
