package p5923

import (
	"strings"
)

//https://leetcode-cn.com/contest/biweekly-contest-66/problems/minimum-number-of-buckets-required-to-collect-rainwater-from-houses/
//https://leetcode-cn.com/problems/minimum-number-of-buckets-required-to-collect-rainwater-from-houses/submissions/
func minimumBuckets(street string) (ans int) {
	n := len(street)
	if street == "." {
		return 0
	}
	if street == "H" {
		return -1
	}
	if strings.Contains(street, "HHH") {
		return -1
	}
	if street[0] == 'H' && street[1] == 'H' {
		return -1
	}
	if street[n-1] == 'H' && street[n-2] == 'H' {
		return -1
	}
	street = strings.Replace(street, "H.H", "hBh", -1)
	street = strings.Replace(street, ".H", "Bh", -1)
	street = strings.Replace(street, "H.", "hB", -1)
	return strings.Count(street, "B")
}
