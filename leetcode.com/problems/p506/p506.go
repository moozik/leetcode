package p506

import (
	"sort"
	"strconv"
)

var gold []string = []string{
	"Gold Medal",
	"Silver Medal",
	"Bronze Medal",
}

type Arr []int

func (t Arr) Len() int {
	return len(t)
}
func (t Arr) Less(i, j int) bool {
	//正序
	// return t[i] < t[j]
	//逆序
	return t[i] > t[j]
}
func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func findRelativeRanks(score []int) (ans []string) {
	var scoreArr Arr = append([]int(nil), score...)
	sort.Sort(scoreArr)
	// fmt.Printf("scoreArr:%+v\n", scoreArr)

	//分数 名次
	goldMap := map[int]int{}
	for ranking, scoreItem := range scoreArr {
		goldMap[scoreItem] = ranking
	}
	for _, scoreItem := range score {
		if ranking := goldMap[scoreItem]; ranking > 2 {
			ans = append(ans, strconv.Itoa(ranking+1))
		} else {
			ans = append(ans, gold[ranking])
		}
	}
	// fmt.Printf("score:%+v\n", score)
	return
}
