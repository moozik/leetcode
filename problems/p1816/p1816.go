package p1816

func truncateSentence(s string, k int) string {
	for index, word := range s {
		if string(word) != " " {
			continue
		}
		k--
		if k == 0 {
			return s[:index]
		}
	}
	return s
}
