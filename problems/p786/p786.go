package p786

//没做出来
func kthSmallestPrimeFraction(arr []int, k int) (ans []int) {
	kk := 1
	n := len(arr)
	posL, posR := 0, n-1
	for kk != k {
		//判断指针移动方向
		if kk+2 <= k {
			posL++
			posR--
			kk++
			continue
		}
		if kk+1 == k {
			if (posL+1)/posR < posL/(posR-1) {
				posL++
			} else {
				posR--
			}
			break
		}
	}
	ans[0], ans[1] = arr[posL], arr[posR]
	return
}
