package leetcode

import . "./lib"

func maxPathSum(root *TreeNode) int {
	result := root.Val
	maxPathSumHelper(root, &result)
	return result
}

func maxPathSumHelper(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}
	l := maxPathSumHelper(node.Left, result)
	r := maxPathSumHelper(node.Right, result)
	if l+r+node.Val > *result {
		*result = l + r + node.Val
	}
	return Max(node.Val+Max(l, r), 0)
}

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
