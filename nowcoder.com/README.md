
## 输入输出总结

单行字符串
```golang
str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
fun(str)
```

指定行字符串
```golang
input :=bufio.NewScanner(os.Stdin) // 和NewReader不同在于scanner读取多行
//读取一行
input.Scan()
str1 := input.Text() // 获取第一行的数据

//读取第二行
input.Scan()
str2 := input.Text()
```


```go
func s2i(s string) int {
    ret,_ := strconv.Atoi(s)
    return ret
}
```


不定行字符串/数字 输入
```golang
func inputTextArr() (ret []string) {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        in := input.Text()
        if in == "" {
            return
        }
        ret = append(ret, in)
    }
    return
}


func inputIntArr() (ret []int) {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        in := input.Text()
        if in == "" {
            return
        }
        intVal,_ := strconv.Atoi(in)
        ret = append(ret, intVal)
    }
    return
}
```

不定行字符串 输出
```golang
func printArr(data []string) {
    for i:=0;i<len(data);i++{
        if data[i] != nil {
            fmt.Println(data[i])
        }
    }
}

func printArr(data []int) {
    for i:=0;i<len(data);i++{
        if data[i] != nil {
            fmt.Println(data[i])
        }
    }
}
```

链表相关
```go
//链表构造
func buildNode(nums []int) (headNode *Node) {
	headNode = &Node{val: nums[0]}
	p := headNode
	for _, item := range nums[1:] {
		p.next = &Node{val: item}
		p = p.next
	}
	return
}

//链表输出
func printNode(node *Node) {
	for {
		fmt.Printf("%d ", node.val)
		if node.next == nil {
			break
		}
		node = node.next
	}
}
```
