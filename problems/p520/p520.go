package p520

func detectCapitalUse(word string) (ans bool) {
	ans = true
	for index, ascii := range word {
		if index == 0 {
			continue
		}
		if word[0] > 91 && ascii < 91 {
			ans = false
			break
		}
		if index > 1 && word[0] < 91 {
			if word[1] < 91 && ascii > 91 {
				ans = false
				break
			}
			if word[1] > 91 && ascii < 91 {
				ans = false
				break
			}
		}
	}
	return
}
