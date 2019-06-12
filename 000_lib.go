package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fill(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
