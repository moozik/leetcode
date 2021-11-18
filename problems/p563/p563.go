package p563

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	return abs(treeSum(root.Left)-treeSum(root.Right)) + findTilt(root.Left) + treeSum(root.Right)
}

func treeSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + treeSum(node.Left) + treeSum(node.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
