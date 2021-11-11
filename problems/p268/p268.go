package p268

//https://leetcode-cn.com/problems/missing-number/

func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return (1+n)*n/2 - sum
}
