package p5941

func findAllPeople(n int, meetings [][]int, firstPerson int) (ans []int) {
	maxTime := 0
	sMap := make(map[int]map[int][]int)
	for _, mee := range meetings {
		maxTime = max(maxTime, mee[2])
		if _, ok := sMap[mee[2]]; !ok {
			sMap[mee[2]] = make(map[int][]int)
		}
		if _, ok := sMap[mee[2]][mee[0]]; !ok {
			sMap[mee[2]][mee[0]] = []int{}
		}
		if _, ok := sMap[mee[2]][mee[1]]; !ok {
			sMap[mee[2]][mee[1]] = []int{}
		}
		sMap[mee[2]][mee[1]] = append(sMap[mee[2]][mee[1]], mee[0])
		sMap[mee[2]][mee[0]] = append(sMap[mee[2]][mee[0]], mee[1])
	}
	ans = append(ans, 0)
	ans = append(ans, firstPerson)
	for i := 0; i <= maxTime; i++ {
		if _, ok := sMap[i]; !ok {
			continue
		}
	}
	return ans
}

func ifIntHave(nums []int, target int) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}

func ifStringHave(nums []string, target string) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
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
