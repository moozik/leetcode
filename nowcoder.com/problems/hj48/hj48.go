package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

/*
6 2 1 2 3 2 5 1 4 5 7 2 2
7 3 1 5 4

5 2 3 2 4 3 5 2 1 4 3
2 5 4 1
*/
func main() {
	var headVal, n int
	fmt.Scan(&n, &headVal)

	nodeHead := &Node{val: headVal}
	for i := 1; i < n; i++ {
		var val, pos int
		fmt.Scan(&val, &pos)

		insert(val, pos, nodeHead)
	}
	// printNode(nodeHead)

	var delVal int
	fmt.Scan(&delVal)

	// fmt.Printf("nodeHead:%v \n", nodeHead)
	nodeHead = delete(delVal, nodeHead)
	// fmt.Printf("nodeHead:%v \n", nodeHead)
	printNode(nodeHead)
}

func printNode(node *Node) {
	for {
		fmt.Printf("%d ", node.val)
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Println()
}

func insert(val, pos int, nodeHead *Node) {
	// fmt.Printf("val:%d,pos:%d \n", val, pos)
	for {
		if nodeHead.val != pos {
			nodeHead = nodeHead.next
			continue
		}
		if nodeHead.next == nil {
			nodeHead.next = &Node{val: val}
			break
		}
		p := Node{val: val, next: nodeHead.next}
		nodeHead.next = &p
		break
	}
}

func delete(val int, headNode *Node) *Node {
	// fmt.Printf("1 val:%d,headNode:%v \n", val, headNode)

	if headNode.val == val {
		return headNode.next
	}

	beforeNode := headNode

	nextNode := beforeNode.next
	for {
		if nextNode.val == val {
			beforeNode.next = nextNode.next
			return headNode
		}
		beforeNode = beforeNode.next
		nextNode = nextNode.next
	}
}
