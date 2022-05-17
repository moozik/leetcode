package p174

//超时解法

var Dungeon [][]int
var xEdge, yEdge int

func calculateMinimumHP(dungeon [][]int) int {
	Dungeon = dungeon
	xEdge = len(dungeon[0]) - 1
	yEdge = len(dungeon) - 1
	minHp := dp(0, 0, 0, 0)
	// fmt.Printf("minHp:%+v\n", minHp)
	if minHp > 0 {
		return 1
	}
	return 0 - minHp + 1
}

func dp(x, y, hp, hpMin int) (retHpMin int) {
	hp += Dungeon[y][x]
	hpMin = min(hpMin, hp)
	if x == xEdge && y == yEdge {
		// fmt.Printf("x:%+v,y:%+v,hp:%+v,hpMin:%+v\n", x, y, hp, hpMin)
		return hpMin
	}
	//往右
	if y == yEdge {
		return dp(x+1, y, hp, hpMin)
	}
	//往下
	if x == xEdge {
		return dp(x, y+1, hp, hpMin)
	}

	return max(dp(x, y+1, hp, hpMin), dp(x+1, y, hp, hpMin))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
