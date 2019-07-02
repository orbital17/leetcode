package leetcode

import . "./lib"

// import "container/list"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	lowestCommonAncestorHelper(root, p, q, &ans)
	return ans
}

func lowestCommonAncestorRec(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	lowestCommonAncestorHelper(root, p, q, &ans)
	return ans
}

func lowestCommonAncestorHelper(node, p, q *TreeNode, ans **TreeNode) bool {
	if node == nil {
		return false
	}
	l := lowestCommonAncestorHelper(node.Left, p, q, ans)
	r := lowestCommonAncestorHelper(node.Right, p, q, ans)
	mid := node == p || node == q
	if l && r || l && mid || r && mid {
		*ans = node
	}
	return l || r || mid
}
