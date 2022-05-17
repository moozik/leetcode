package p5924

//https://leetcode-cn.com/contest/biweekly-contest-66/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/
//https://leetcode-cn.com/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) (ans int) {
	//已经在家
	if startPos[0] == homePos[0] && startPos[1] == homePos[1] {
		return 0
	}
	var rowHomeDirection, colHomeDirection int
	rowNow := startPos[0]
	colNow := startPos[1]

	if startPos[0] == homePos[0] {
		rowHomeDirection = 0
	} else if startPos[0] > homePos[0] {
		rowHomeDirection = -1
	} else {
		rowHomeDirection = 1
	}
	if startPos[1] == homePos[1] {
		colHomeDirection = 0
	} else if startPos[1] > homePos[1] {
		colHomeDirection = -1
	} else {
		colHomeDirection = 1
	}

	if rowHomeDirection != 0 {
		for rowNow != homePos[0] {
			rowNow += rowHomeDirection
			ans += rowCosts[rowNow]
		}
	}
	if colHomeDirection != 0 {
		for colNow != homePos[1] {
			colNow += colHomeDirection
			ans += colCosts[colNow]
		}
	}
	return
}
