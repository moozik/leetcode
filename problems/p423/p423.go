package p423

import "strings"

//"e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"

//zero one two three four five sex seven eight nine

//(z)ero t(w)o fo(u)r se(x) ei(g)ht

//(o)ne (t)hree (f)ive (s)even nine

func originalDigits(s string) string {
	numC := map[byte]int{}
	for _, item := range s {
		numC[byte(item)]++
	}
	//0
	if co, ok := numC['z']; ok && co != 0 {
		numC['o'] -= co
	}
	//2
	if co, ok := numC['w']; ok && co != 0 {
		numC['t'] -= co
		numC['o'] -= co
	}
	//4
	if co, ok := numC['u']; ok && co != 0 {
		numC['f'] -= co
		numC['o'] -= co
	}
	//6
	if co, ok := numC['x']; ok && co != 0 {
		numC['s'] -= co
	}
	//8
	if co, ok := numC['g']; ok && co != 0 {
		numC['t'] -= co
	}
	//1
	if co, ok := numC['o']; ok && co != 0 {
		numC['n'] -= co
	}
	//3
	// if co, ok := numC['t']; ok && co != 0 {
	// }
	//5
	// if co, ok := numC['f']; ok && co != 0 {
	// }
	//7
	if co, ok := numC['s']; ok && co != 0 {
		numC['n'] -= co
	}
	//9
	// if co, ok := numC['n']; ok && co != 0 {
	// }
	ret := ""
	ret += strings.Repeat("0", numC['z'])
	ret += strings.Repeat("1", numC['o'])
	ret += strings.Repeat("2", numC['w'])
	ret += strings.Repeat("3", numC['t'])
	ret += strings.Repeat("4", numC['u'])
	ret += strings.Repeat("5", numC['f'])
	ret += strings.Repeat("6", numC['x'])
	ret += strings.Repeat("7", numC['s'])
	ret += strings.Repeat("8", numC['g'])
	ret += strings.Repeat("9", numC['n']/2)
	return ret
}
