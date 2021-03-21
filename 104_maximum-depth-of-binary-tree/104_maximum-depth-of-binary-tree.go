package leetcode

import . "github.com/orbital17/leetcode/lib"

func maxDepth(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepthDFS(root *TreeNode) (res int) {
	var old, new []*TreeNode
	old = append(old, root)
	for len(old) != 0 {
		for _, n := range old {
			if n != nil {
				new = append(new, n.Left)
				new = append(new, n.Right)
			}
		}
		old = new
		new = []*TreeNode{}
		if len(old) != 0 {
			res++
		}
	}
	return
}
