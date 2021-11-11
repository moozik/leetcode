package p495

func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
	ans = duration
	n := len(timeSeries)
	for i := 1; i < n; i++ {
		//延长中毒时间
		ans += duration
		//修正中毒时间
		if timeSeries[i]-timeSeries[i-1] < duration {
			ans -= duration - (timeSeries[i] - timeSeries[i-1])
		}
	}
	return
}
