package p174_2

func calculateMinimumHP(dungeon [][]int) int {
	yLen := len(dungeon)
	xLen := len(dungeon[0])
	dungeonPath := make([][]int, yLen)
	for y := range dungeonPath {
		dungeonPath[y] = make([]int, xLen)
	}
	return 0
}
