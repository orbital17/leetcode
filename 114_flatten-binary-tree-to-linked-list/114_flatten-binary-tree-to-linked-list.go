package leetcode

import . "github.com/orbital17/leetcode/lib"

func flatten(root *TreeNode) {
	var cur, a *TreeNode
	for cur = root; cur != nil; cur = cur.Right {
		if cur.Left != nil {
			a = cur.Left
			for a.Right != nil {
				a = a.Right
			}
			a.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
	}
}
