package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	list = make([]string, 0)
	traversal(root, "")
	sort.Strings(list)

	return list[0]
}

var list []string

func traversal(node *TreeNode, str string) {
	if node == nil {
		return
	}

	str = string(node.Val+97) + str

	if node.Left == nil && node.Right == nil {
		list = append(list, str)
		return
	}

	traversal(node.Left, str)
	traversal(node.Right, str)
}
