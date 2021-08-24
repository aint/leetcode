package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var values []int

func findTarget(root *TreeNode, k int) bool {
	values = make([]int, 0)
	collectValues(root)

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == k {
				return true
			}
		}
	}

	return false
}

func collectValues(root *TreeNode) {
	if root == nil {
		return
	}

	values = append(values, root.Val)

	collectValues(root.Left)
	collectValues(root.Right)
}
