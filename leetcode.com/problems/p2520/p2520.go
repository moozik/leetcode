package p2520

import "strconv"

func countDigits(num int) int {
	if num < 10 {
		return 1
	}
	ret := 0
	for _, s := range strconv.Itoa(num) {
		if num%(int(s)-48) == 0 {
			ret++
		}
	}
	return ret
}
