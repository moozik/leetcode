package p458

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) (ans int) {
	ans = 1
	if minutesToDie == minutesToTest {
		for ; ; ans++ {
			if math.Pow(2, float64(ans)) > float64(buckets) {
				return ans
			}
		}
	}
	return
}
