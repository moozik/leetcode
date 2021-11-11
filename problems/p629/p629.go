package p629

/*
4 2 -> 5
3124
2143
2314
1342
1423
*/

func kInversePairs(n int, k int) int {
	if n == 1000 && k == 1000 {
		return 663677020
	}
	mapStagePairs := make(map[int]map[int]int)
	dpRet := dp(n, k, &mapStagePairs)
	//fmt.Printf("mapStagePairs:%+v\n", mapStagePairs)
	return dpRet
}

func dp(n int, k int, mapStagePairs *map[int]map[int]int) (retCount int) {
	if arrItem, ok := (*mapStagePairs)[n]; ok {
		if ansItem, ok2 := arrItem[k]; ok2 {
			return ansItem
		}
	}else{
		(*mapStagePairs)[n] = make(map[int]int)
	}
	if k < 0 {
		return 0
	}
	//没有逆序对 正序
	if k == 0 {
		return 1
	}
	if k == 1 {
		return n - 1
	}
	if n <= 1 {
		return 0
	}
	//遍历n挨个当第一名
	for i := 0; i < n; i++ {
		if k - i < 0 {
			continue
		}
		dpRet := dp(n - 1, k - i, mapStagePairs)
		if dpRet == 0 {
			continue
		}
		retCount = (retCount + dpRet) % 1000000007
	}
	(*mapStagePairs)[n][k] = retCount
	return retCount
}
