package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const A = 0
const B = 1
const C = 2
const D = 3
const E = 4
const WRONG = 5
const PRIVATE = 6

const IP = `^((\d|1?\d\d|2[0-4]\d|25[0-5])\.){3}(\d|1?\d\d|2[0-4]\d|25[0-5])$`

func main() {
	ret := fun(inputTextArr())
	fmt.Println(ret[0], ret[1], ret[2], ret[3], ret[4], ret[5], ret[6])
}

func inputTextArr() []string {
	input := bufio.NewScanner(os.Stdin)
	data := make([]string, 0)
	for input.Scan() {
		in := input.Text()
		if in == "" {
			return data
		}
		data = append(data, in)
	}
	return data
}

func ifMaskRight(mask string) bool {
	if ok, _ := regexp.MatchString(IP, mask); !ok {
		//ip错误
		return false
	}
	r, _ := regexp.Compile(`^(255\.255\.255\.|255\.255\.|255\.)`)
	mask = r.ReplaceAllString(mask, "")
	//fmt.Println("mask1", mask)

	r, _ = regexp.Compile(`(\.0)+`)
	mask = r.ReplaceAllString(mask, "")
	//fmt.Println("mask2", mask)
	if mask == "" || mask == "0" {
		return true
	}
	if !isInt(mask) {
		return false
	}
	maskInt, _ := strconv.Atoi(mask)
	return intHave([]int{128, 192, 224, 240, 248, 252, 254}, maskInt)
}

/*

私网IP范围是：

从10.0.0.0到10.255.255.255

从172.16.0.0到172.31.255.255

从192.168.0.0到192.168.255.255
*/
func isPrivatIP(ip1 int, ip2 int) bool {
	if ip1 == 10 {
		return true
	}
	if ip1 == 197 && ip2 >= 16 && ip2 <= 31 {
		return true
	}
	if ip1 == 192 && ip2 == 168 {
		return true
	}
	return false
}

func intHave(nums []int, target int) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func fun(strArr []string) (ret []int) {
	ret = make([]int, 7)
	for _, item := range strArr {
		ipMask := strings.Split(item, "~")
		if ipMask[0][:2] == "0." || ipMask[0][:4] == "127." {
			//fmt.Println("忽略", item)
			//忽略
			continue
		}
		if ok, _ := regexp.MatchString(IP, ipMask[0]); !ok {
			//fmt.Println("ip错误", item)
			//ip错误
			ret[WRONG]++
			continue
		}
		if !ifMaskRight(ipMask[1]) {
			//fmt.Println("mask错误", item)
			//mask错误
			ret[WRONG]++
			continue
		}
		ipArr := strings.Split(ipMask[0], ".")
		ipInt, _ := strconv.Atoi(ipArr[0])
		ipInt2, _ := strconv.Atoi(ipArr[1])

		if isPrivatIP(ipInt, ipInt2) {
			//fmt.Println("私有ip", ipMask[0])
			ret[PRIVATE]++
		}

		if ipInt >= 1 && ipInt <= 126 {
			ret[A]++
			continue
		}
		if ipInt >= 128 && ipInt <= 191 {
			ret[B]++
			continue
		}
		if ipInt >= 192 && ipInt <= 223 {
			ret[C]++
			continue
		}
		if ipInt >= 224 && ipInt <= 239 {
			ret[D]++
			continue
		}
		if ipInt >= 240 && ipInt <= 255 {
			ret[E]++
			continue
		}
	}
	return
}
