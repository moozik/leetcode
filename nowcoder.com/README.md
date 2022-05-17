
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
        ret = append(ret, strconv.Atoi(in))
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