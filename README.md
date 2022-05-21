# leetcode

## 常用函数
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


## 数组操作
```golang
//排序
type Arr []int

func (t Arr) Len() int {
	return len(t)
}
func (t Arr) Less(i, j int) bool {
	//正序
	return t[i] < t[j]
	//逆序
	//return t[i] > t[j]
}
func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
```

字符串函数
```golang
//字符串计数
strings.Count(row, "O")
```

```golang
func intHave(nums []int, target int) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}

func stringHave(nums []string, target string) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}
```


```go
func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}


func isInt(s string) bool {
	_, err := strconv.ParseInt(s)
	return err == nil
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
```