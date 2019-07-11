package leetcode

import . "../lib"

func rob(root *TreeNode) int {
	return max(recHelper(root))
}

func recHelper(node *TreeNode) (broken, unbroken int) {
	if node == nil {
		return 0, 0
	}
	lb, lu := recHelper(node.Left)
	rb, ru := recHelper(node.Right)
	broken = node.Val + lu + ru
	unbroken = max(lb, lu) + max(rb, ru)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
