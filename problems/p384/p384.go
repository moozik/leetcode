package p384

type Solution struct {
	n          int
	numsOrigin []int
	nums       []int
}

func Constructor(nums []int) Solution {
	return Solution{numsOrigin: nums, n: len(nums)}
}

func (t *Solution) Reset() []int {
	return t.numsOrigin
}

func (t *Solution) Shuffle() []int {
	t.nums = t.numsOrigin
	//洗牌算法
	for i := 0; i < t.n-1; i++ {

	}
	return t.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
