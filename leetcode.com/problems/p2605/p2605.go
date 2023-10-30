package p2605

func minNumber(nums1 []int, nums2 []int) int {
	flag := make([]int8, 10, 10)
	flag[0] = 10
	for _, i := range nums1 {
		flag[i] = 1
		nums1[0] = min(nums1[0], i)
	}
	for _, i := range nums2 {
		if flag[i] == 1 {
			flag[0] = min(flag[0], int8(i))
		}
		nums2[0] = min(nums2[0], i)
	}
	if flag[0] < 10 {
		return int(flag[0])
	}
	return min(nums1[0], nums2[0])*10 + max(nums1[0], nums2[0])
}

func min[T int8 | int32 | int64 | int](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T int8 | int32 | int64 | int](x, y T) T {
	if x > y {
		return x
	}
	return y
}
