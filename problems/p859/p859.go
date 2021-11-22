package p859

func buddyStrings(s string, goal string) bool {
	sn := len(s)
	if sn < 2 {
		return false
	}
	if sn != len(goal) {
		return false
	}
	//字符串一致
	if s == goal {
		if sn > 26 {
			return true
		}
		for i := 0; i < sn-1; i++ {
			for o := i + 1; o < sn; o++ {
				if s[i] == s[o] {
					return true
				}
			}
		}
	}
	//字符串不一致
	first, second := -1, -1
	for i := 0; i < sn; i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return first != -1 && second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
