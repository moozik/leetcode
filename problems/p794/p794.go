package p794

func validTicTacToe(board []string) bool {
	pointCount := make(map[string]int, 3)

	for i := 0; i < 3; i++ {
		for o := 0; o < 3; o++ {
			pointCount[string(board[i][o])]++
		}
	}
	// fmt.Printf("pointCount:%+v\n", pointCount)
	cha := pointCount["X"] - pointCount["O"]
	//个数不支持
	if cha != 0 && cha != 1 {
		return false
	}
	if win(board, 'X') && win(board, 'O') {
		return false
	}
	if cha == 0 && win(board, 'X') {
		return false
	}
	if cha == 1 && win(board, 'O') {
		return false
	}
	return true
}

func win(board []string, b byte) bool {
	for i := 0; i < 3; i++ {
		//判断横竖
		if board[i][0] == b && board[i][1] == b && board[i][2] == b {
			return true
		}
		if board[0][i] == b && board[1][i] == b && board[2][i] == b {
			return true
		}
	}
	if board[0][0] == b && board[1][1] == b && board[2][2] == b {
		return true
	}
	if board[0][2] == b && board[1][1] == b && board[2][0] == b {
		return true
	}
	return false
}
