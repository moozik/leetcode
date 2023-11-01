package p849

import (
	"math"
)

func maxDistToClosest(seats []int) int {
	p1 := 0
	p2 := 1
	m := 0
	for p2 < len(seats) {
		if seats[p2] == 1 {
			if seats[p1] == 0 {
				m = Max(m, p2)
			} else {
				m = Max(m, int(math.Floor((float64(p2-p1))/2)))
			}
			p1 = p2
		}
		p2++
	}
	if seats[len(seats)-1] == 0 {
		m = Max(m, len(seats)-1-p1)
	}
	return m
}

func Max[T int8 | int32 | int64 | int | float32 | float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}
