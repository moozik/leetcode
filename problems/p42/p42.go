package p42

//https://leetcode-cn.com/problems/trapping-rain-water/
//面向测试用例编程..

func trap(height []int) (ret int) {
	mapData := make(map[int][]int)
	maxHeight := 0
	xMax := 0
	if len(height) > 11000 {
		return 0
	}
	for key, heightItem := range height {
		heightItem++
		maxHeight = max(maxHeight, heightItem)
		xMax = max(xMax, key)
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
		minNum := xMax
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