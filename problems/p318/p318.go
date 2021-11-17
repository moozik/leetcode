package p318

func maxProduct(words []string) (ans int) {
	wordsBitArr := make([]int, len(words))
	for index, word := range words {
		wordsBitArr[index] = 0
		for _, ch := range word {
			//1. 数组默认值为0
			//2. 1 << (ch - 'a'),代表将1向左移动0-25位
			//3. |=是求且，将每一位的比特值合并
			wordsBitArr[index] |= 1 << (ch - 'a')
		}
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if wordsBitArr[i]&wordsBitArr[j] == 0 {
				tmp := len(words[i]) * len(words[j])
				if tmp > ans {
					ans = tmp
				}
			}
		}
	}
	return
}
