//超时了
package p5186

type RangeFreqQuery struct {
	MapStorage map[int]map[int]int
}

func Constructor(arr []int) RangeFreqQuery {
	ret := RangeFreqQuery{
		MapStorage: make(map[int]map[int]int),
	}
	for index, num := range arr {
		if _, ok := ret.MapStorage[num]; !ok {
			ret.MapStorage[num] = make(map[int]int)
			ret.MapStorage[num][-1] = 0
		}
		ret.MapStorage[num][-1]++
		ret.MapStorage[num][index] = ret.MapStorage[num][-1]
	}
	// fmt.Printf("ret.MapStorage:%+v\n", ret.MapStorage)
	return ret
}

func (t *RangeFreqQuery) Query(left int, right int, value int) (ans int) {
	ear := 0
	las := 0
	for i := left; i <= right; i++ {
		if co, ok := t.MapStorage[value][i]; ok {
			if ear == 0 {
				ear = co
			}
			las = co
		}
	}
	if ear == 0 && las == 0 {
		return 0
	}
	if ear == las {
		return 1
	}
	return las - ear + 1
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
