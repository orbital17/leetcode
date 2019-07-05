package leetcode

import . "../lib"

func levelOrder(root *TreeNode) (res [][]int) {
	var old, new []*TreeNode
	old = append(old, root)
	cur := []int{}
	for len(old) != 0 {
		for _, n := range old {
			if n != nil {
				cur = append(cur, n.Val)
				new = append(new, n.Left)
				new = append(new, n.Right)
			}
		}
		old = new
		new = []*TreeNode{}
		if len(cur) != 0 {
			res = append(res, cur)
			cur = []int{}
		}
	}
	return
}
