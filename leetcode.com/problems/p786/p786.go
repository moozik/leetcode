package p786

//https://leetcode-cn.com/problems/k-th-smallest-prime-fraction/
import (
	"sort"
)

type Arr [][]int

func (t Arr) Len() int {
	return len(t)
}
func (t Arr) Less(i, j int) bool {
	return t[i][0]*t[j][1] < t[j][0]*t[i][1]
}
func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

//通过排序来确定
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	var arrSort Arr
	for i := 0; i < n-1; i++ {
		for o := i + 1; o < n; o++ {
			arrSort = append(arrSort, []int{arr[i], arr[o]})
		}
	}
	sort.Sort(arrSort)
	return arrSort[k-1]
}
