package leetcode

func inorderTraversalRecursive(root *TreeNode) []int {
	result := []int{}
	inorderTraversal_(&result, root)
	return result
}

func inorderTraversal_(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraversal_(res, node.Left)
	*res = append(*res, node.Val)
	inorderTraversal_(res, node.Right)
}
