/**
// 给你一个非递减数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 非递减：nums[j] >= nums[i], j > i
// 要求时间复杂度为 O(n)

// nums = [1,2,3,4,5]
// sorted(nums) = [1,4,9,16,25]

// nums = [-4,-1,0,3,10]
// sorted(nums) = [0,1,9,16,100]

// nums = [-7,-3,2,3,11]
// sorted(nums) = [4,9,9,49,121]

package main

import(
  "fmt"
)
func sorted(nums []int) []int {
    // TODO 实现方法
    return nil
}

func main() {
    cases := []*struct {
        nums []int
        want []int
    }{
        {
            nums: []int{1, 2, 3, 4, 5},//1
            want: []int{1, 4, 9, 16, 25},
        },
        {
            nums: []int{-4, -1, 0, 3, 10},//0
            want: []int{0, 1, 9, 16, 100},
        },
        {
            nums: []int{-7, -3, 2, 3, 11},//2
            want: []int{4, 9, 9, 49, 121},
        },
    }

    for i, c := range cases {
        got := sorted(c.nums)
        if !reflect.DeepEqual(got, c.want) {
            fmt.Printf("case %d want %v but got %v\n", i, c.want, got)
        }
    }
}

*/
package main

import (
	"fmt"
	"reflect"
)

func sorted(nums []int) (ret []int) {
	var pindex = 0
	for k, item := range nums {
		pindex = k
		if item >= 0 {
			break
		}
	}
	// fmt.Println("pindex", pindex)
	p1, p2 := pindex, pindex+1
	for p1 >= 0 || p2 < len(nums) {
		p1num, p2num := 0, 0
		if p1 >= 0 {
			p1num = nums[p1] * nums[p1]
		}
		if p2 < len(nums) {
			p2num = nums[p2] * nums[p2]
		}
		if p1 >= 0 && p2 < len(nums) && p1num < p2num {
			ret = append(ret, p1num)
			p1--
			continue
		}
		if p1 >= 0 && p2 < len(nums) && p1num >= p2num {
			ret = append(ret, p2num)
			p2++
			continue
		}

		if p1 >= 0 {
			ret = append(ret, p1num)
			p1--
			continue
		}

		if p2 < len(nums) {
			ret = append(ret, p2num)
			p2++
			continue
		}
	}
	return
}

func main() {
	cases := []*struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			nums: []int{-17, -13, -5, -4, -2},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for i, c := range cases {
		got := sorted(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			fmt.Printf("case %d want %v but got %v\n", i, c.want, got)
		}
	}
}
