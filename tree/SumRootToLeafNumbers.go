package main

import (
	"strconv"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum = 0
	dfs(root, "")
	return sum
}

var sum int

func dfs(node *TreeNode, numStr string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		numStr += strconv.Itoa(node.Val)
		n, _ := strconv.Atoi(numStr)
		sum += n
		return
	}

	dfs(node.Left, numStr+strconv.Itoa(node.Val))
	dfs(node.Right, numStr+strconv.Itoa(node.Val))
}
