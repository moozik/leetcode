package p42

//https://leetcode-cn.com/problems/trapping-rain-water/
//面向测试用例编程..

func trap(height []int) (ret int) {
	mapData := make(map[int][]int)
	xLeft := -1
	xRight := -1
	//用单调性剪枝
	for key, hItem := range height {
		if key != 0 {
			if xLeft == -1 && hItem < height[key-1] {
				xLeft = key - 1
				continue
			}
			if xLeft != -1 {
				if hItem > height[key-1] {
					xRight = max(xRight, key)
				}
			}
		}
	}
	//fmt.Printf("xLeft:%+v,xRight:%+v\n",xLeft, xRight)
	if xLeft < 0 || xRight < 0 {
		return -1
	}

	maxHeight := -1
	for key, heightItem := range height {
		if key < xLeft || key > xRight {
			continue
		}
		heightItem++
		maxHeight = max(maxHeight, heightItem)
		if _, ok := mapData[heightItem]; ok {
			mapData[heightItem] = append(mapData[heightItem], key)
			continue
		}
		mapData[heightItem] = []int{key}
	}

	//补充数据
	if maxHeight >= 2 {
		for i := maxHeight - 1; i >= 1; i-- {
			if _, ok := mapData[i]; ok {
				mapData[i] = append(mapData[i], mapData[i+1]...)
			} else {
				mapData[i] = mapData[i+1]
			}
		}
	}
	//fmt.Printf("mapData:%+v\n", mapData)
	//一层一层切掉
	for level := maxHeight; level >= 1; level-- {
		if len(mapData[level]) == 1 {
			continue
		}
		if len(mapData[level]) == 2 {
			ret += abs(mapData[level][0]-mapData[level][1]) - 1
			continue
		}
		maxNum := -1
		minNum := xRight
		for _, item := range mapData[level] {
			maxNum = max(maxNum, item)
			minNum = min(minNum, item)
		}
		ret += (maxNum - minNum - 1) - (len(mapData[level]) - 2)
	}
	return ret
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
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
