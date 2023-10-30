package p1465

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH := 1
	lastH := 0
	for _, h = range horizontalCuts {
		if h-lastH > maxH {
			maxH = h - lastH
		}
		lastH = h
	}
	maxW := 1
	lastW := 0
	for _, w = range verticalCuts {
		if w-lastW > maxW {
			maxW = w - lastW
		}
		lastW = w
	}
	//fmt.Println(maxH, maxW, maxH*maxW)
	return (maxH * maxW) % 1000000007
}
