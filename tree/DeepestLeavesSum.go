package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]int

func deepestLeavesSum(root *TreeNode) int {
	m = make(map[int]int)
	dfs(root, 0)

	max, res := 0, 0
	for k, v := range m {
		if k >= max {
			res = v
			max = k
		}
	}
	return res
}

func dfs(root *TreeNode, count int) {
	if root == nil {
		return
	}

	m[count] += root.Val

	count += 1
	dfs(root.Left, count)
	dfs(root.Right, count)
}
