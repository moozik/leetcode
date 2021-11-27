package p5922

//https://leetcode-cn.com/contest/biweekly-contest-66/problems/count-common-words-with-one-occurrence/
func countWords(words1 []string, words2 []string) (ans int) {
	wordMap1, wordMap2 := map[string]int{}, map[string]int{}
	for _, word := range words1 {
		wordMap1[word]++
	}
	for _, word := range words2 {
		wordMap2[word]++
	}
	for word, count := range wordMap1 {
		if count != 1 {
			continue
		}
		if count2, ok := wordMap2[word]; ok && count2 == 1 {
			ans++
		}
	}
	return
}
