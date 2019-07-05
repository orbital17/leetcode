package leetcode

import . "../lib"

func inorderTraversal(root *TreeNode) []int {
	stack := TreeStack{}
	cur := root
	result := []int{}
	for cur != nil || !stack.IsEmpty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Pop()
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}
	return result
}

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
