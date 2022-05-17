package p384

import (
	"math/rand"
)

type Solution struct {
	n          int
	numsOrigin []int
	nums       []int
}

func Constructor(nums []int) Solution {
	return Solution{numsOrigin: nums, nums: append([]int(nil), nums...), n: len(nums)}
}

func (t *Solution) Reset() []int {
	copy(t.nums, t.numsOrigin)
	return t.nums
}

func (t *Solution) Shuffle() []int {
	//洗牌算法
	for i := 0; i < t.n-1; i++ {
		randI := i + rand.Intn(t.n-i)
		t.nums[i], t.nums[randI] = t.nums[randI], t.nums[i]
	}
	return t.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
