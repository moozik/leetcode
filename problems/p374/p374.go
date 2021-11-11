package p374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

//guess 模拟guess
func guess(n, num int) int {
	ans := (n*69+857092)%n + 1
	if ans > num {
		return 1
	}
	if ans < num {
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	if n == 1 {
		return 1
	}
	leftPos := 1
	rightPos := n
	guessNum := n >> 1
	for {
		if leftPos == rightPos {
			break
		}
		//需要把n去掉提交
		guessRet := guess(n, guessNum)
		if guessRet == 0 {
			break
		}
		if guessRet == -1 {
			rightPos = guessNum
		}
		if guessRet == 1 {
			leftPos = guessNum + 1
		}
		guessNum = (rightPos + leftPos) >> 1
	}
	return guessNum
}
