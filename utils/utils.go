package utils

import (
	"strconv"
	"strings"
)

func Str2int(str string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(str))
	return i
}

//IntArr [3,0,8,4] 转换 为数组
func IntArr(str string) (ret []int) {
	str = strings.Trim(str, "[]")
	for _, item := range strings.Split(str, ",") {
		ret = append(ret, Str2int(item))
	}
	return
}

//GenIntArray [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]] 转换 为数组
func GenIntArray(str string) (ret [][]int) {
	str = strings.Trim(str, "[]")
	intStrArr := strings.Split(str, "],[")
	ret = make([][]int, len(intStrArr))
	for key, strArr := range intStrArr {
		arr := strings.Split(strArr, ",")
		ret[key] = make([]int, len(arr))
		for key2, intStr := range arr {
			ret[key][key2] = Str2int(intStr)
		}
	}
	return
}
