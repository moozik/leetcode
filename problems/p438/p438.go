package p438

//https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) (ans []int) {
	nP := len(p)
	nS := len(s)
	if nP > nS {
		return []int{}
	}
	mapS, mapP := make(map[byte]int), make(map[byte]int)
	for _, ch := range p {
		mapP[byte(ch)]++
	}
	for i := 0; i < nP-1; i++ {
		mapS[s[i]]++
	}
	for i := 0; i <= nS-nP; i++ {
		mapS[s[i+nP-1]]++
		if i > 0 {
			mapS[s[i-1]]--
		}
		flag := true
		for ch, countP := range mapP {
			if countS, ok := mapS[ch]; !ok || countP != countS {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}
