package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

//多组输入
/*
1
8108
1
7
2542 4218 9064 4908 1526 6655 921
1
*/
func main() {
	var n, target int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&nums[i])
		}
		fmt.Scan(&target)
		headNode := buildNode(nums)
		ret := fun(target, headNode)
		fmt.Println(ret)
	}
}

func buildNode(nums []int) (headNode *Node) {
	headNode = &Node{val: nums[0]}
	p := headNode
	for _, item := range nums[1:] {
		p.next = &Node{val: item}
		p = p.next
	}
	return
}

func fun(target int, headNode *Node) int {
	var p1, p2 *Node
	p1 = headNode
	p2 = headNode
	for i := 0; i < target-1; i++ {
		p2 = p2.next
	}
	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p1.val
}
