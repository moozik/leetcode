package p1103

import "fmt"

func distributeCandiesOld(candies int, num_people int) []int {
	ret := make([]int, num_people)
	plus := 0
	for i := 0; candies != 0; i++ {
		if i == num_people {
			i = 0
		}
		plus += 1
		if plus <= candies {
			candies -= plus
			ret[i] += plus
			fmt.Println(ret)
			continue
		}
		ret[i] += candies
		candies = 0
		fmt.Println(ret)
		break
	}
	return ret
}

func distributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)

	for i := 0; i < num_people; i++ {
		if candies >= i+1 {
			ret[i] = i + 1
			candies -= i + 1
			continue
		}
		ret[i] = candies
		return ret
	}
	fmt.Println(ret)

	rowValue := num_people * num_people
	loopCount := candies / rowValue
	if loopCount == 0 {
		for i := 0; i < num_people; i++ {
			if candies >= i+1+num_people {
				ret[i] += i + 1 + num_people
				candies -= i + 1 + num_people
				continue
			}
			ret[i] += candies
			return ret
		}
	}
	fmt.Println("loopCount", loopCount)
	for i := 0; i < num_people; i++ {
		ret[i] += loopCount * num_people
		candies -= loopCount * num_people
	}
	fmt.Println(ret)

	loopTail := candies - (loopCount * rowValue)
	plus := ret[num_people-1]
	for i := 0; i < num_people; i++ {
		plus += 1
		if plus <= loopTail {
			loopTail -= plus
			ret[i] += plus
			continue
		} else {
			ret[i] += loopTail
			break
		}
	}
	fmt.Println(ret)
	return ret
}
