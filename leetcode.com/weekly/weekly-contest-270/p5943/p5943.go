package p5943

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	nodeLen := 1
	nodeHead := head
	nodeNow := head
	for {
		if nodeNow.Next == nil {
			break
		}
		nodeNow = nodeNow.Next
		nodeLen++
	}
	if nodeLen == 1 {
		return nil
	}

	nodeMid := nodeLen >> 1
	nodeNow = head
	lastNode := head
	nodePos := 0
	for {
		if nodePos == nodeMid {
			lastNode.Next = nodeNow.Next
			break
		}
		lastNode = nodeNow
		nodeNow = nodeNow.Next
		nodePos++
	}
	return nodeHead
}
