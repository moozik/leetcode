package offer069

import "fmt"

func peakIndexInMountainArray(arr []int) (retPos int) {
	n := len(arr)
	posLeft := 0
	posRight := n - 1
	var posMid int
	for i := 0; i < 100; i++ {
		fmt.Printf("posLeft:%+v,posRight:%+v\n", posLeft, posRight)
		//返回点
		if posLeft == posRight || abs(posLeft-posRight) == 1 {
			if arr[posLeft] > arr[posRight] {
				return posLeft
			}
			return posRight
		}
		//看中点单调性
		posMid = (posRight-posLeft)>>1 + posLeft
		if arr[posMid-1] > arr[posMid] {
			//取左边
			posRight = posMid - 1
			continue
		}
		if arr[posMid+1] > arr[posMid] {
			//取右边
			posLeft = posMid + 1
			continue
		}
		return posMid
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
