package leetcode

import . "./lib"

func isSymmetricTraversal(root *TreeNode) bool {
	if root == nil {
		return true
	}
	a, b := root.Left, root.Right
	ls, rs := StackAbstract{}, StackAbstract{}
	for a != nil || b != nil || !ls.IsEmpty() || !rs.IsEmpty() {
		switch {
		case a != nil && b != nil:
			if a.Val != b.Val {
				return false
			}
			ls.Push(a)
			rs.Push(b)
			a = a.Left
			b = b.Right
		case a == nil && b == nil:
			if ls.IsEmpty() || rs.IsEmpty() {
				return false
			}
			a = ls.Pop().(*TreeNode)
			b = rs.Pop().(*TreeNode)
			a = a.Right
			b = b.Left
		default:
			return false
		}
	}
	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirroredTrees(root.Left, root.Right)
}

func isMirroredTrees(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	res := isMirroredTrees(a.Left, b.Right) &&
		isMirroredTrees(a.Right, b.Left)
	return res
}
