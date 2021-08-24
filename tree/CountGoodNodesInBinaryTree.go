package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func goodNodes(root *TreeNode) int {
	count = 0
	rec(root, root.Val)
	return count
}

func rec(root *TreeNode, max int) {
	if root == nil {
		return
	}

	if root.Val >= max {
		max = root.Val
		count += 1
	}

	rec(root.Left, max)
	rec(root.Right, max)
}
