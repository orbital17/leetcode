package leetcode

import . "../lib"

func isValidBST(root *TreeNode) bool {

	stack := TreeStack{}
	cur := root
	prevInit := false
	prev := 0
	for cur != nil || !stack.IsEmpty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Pop()
			if prevInit {
				if cur.Val <= prev {
					return false
				}
			} else {
				prevInit = true
			}
			prev = cur.Val
			cur = cur.Right
		}
	}
	return true
}
