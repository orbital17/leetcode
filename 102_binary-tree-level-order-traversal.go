package leetcode

import . "./lib"

func levelOrder(root *TreeNode) (res [][]int) {
	var old, new Queue
	old.Add(root)
	cur := []int{}
	for !old.IsEmpty() {
		for !old.IsEmpty() {
			n := old.Poll().(*TreeNode)
			if n != nil {
				cur = append(cur, n.Val)
				new.Add(n.Left)
				new.Add(n.Right)
			}
		}
		old = new
		new = Queue{}
		if len(cur) != 0 {
			res = append(res, cur)
			cur = []int{}
		}
	}
	return
}
