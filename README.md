# leetcode

常用函数
```golang

//string2int64
i,err := strconv.ParseInt(string, 10, 64)

//int642string
s := strconv.FormatInt(int64,10)

//string2int
i, err := strconv.Atoi(string)

//int2string
s := strconv.Itoa(int)

//int2float32
float := float32(int)

//int2float64
float := float64(int)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```
