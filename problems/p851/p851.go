package p851

//https://leetcode-cn.com/problems/loud-and-rich/
//超时解法
func loudAndRich(richer [][]int, quiet []int) (ans []int) {
	n := len(quiet)
	for person := 0; person < n; person++ {
		personQuietMin := person
		// fmt.Printf("richer:%+v\n", richer)
		for _, richetPerson := range getRicherArr(richer, person) {
			// fmt.Printf("person:%+v,richetPerson:%+v\n", person, richetPerson)
			if quiet[richetPerson] < quiet[personQuietMin] {
				personQuietMin = richetPerson
			}
		}
		ans = append(ans, personQuietMin)
	}
	return
}

func getRicherArr(richerOri [][]int, person int) (ans []int) {
	ans = append(ans, person)
	findRicher := false
	richer := make([]*[]int, len(richerOri))
	for index := 0; index < len(richerOri); index++ {
		richer[index] = &richerOri[index]
	}
	for {
		findRicher = false

		for index := 0; index < len(richer); {
			item := richer[index]
			// fmt.Printf("index:%+v, richer:%+v, item:%+v\n", index, richer, (*item))
			if !intHave(ans, (*item)[0]) && intHave(ans, (*item)[1]) {
				ans = append(ans, (*item)[0])
				findRicher = true
				richer = append(richer[:index], richer[index+1:]...)
				continue
			}
			index++
		}
		if !findRicher {
			break
		}
	}
	// fmt.Printf("ans:%+v,person:%+v\n", ans, person)
	return
}

func intHave(nums []int, target int) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}
