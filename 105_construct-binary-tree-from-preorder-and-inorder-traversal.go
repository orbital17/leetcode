package leetcode

import . "./lib"

type buildTreeInfo struct {
	in, pre, len int
	node         *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	queue := Queue{}
	queue.Add(buildTreeInfo{
		0, 0, len(preorder), root,
	})
	var lleft, lright int
	for !queue.IsEmpty() {
		info := queue.Poll().(buildTreeInfo)
		info.node.Val = preorder[info.pre]
		for i := info.in; i < info.in+info.len; i++ {
			if inorder[i] == preorder[info.pre] {
				lleft = i - info.in
				lright = info.len - i - 1 + info.in
				if lleft > 0 {
					info.node.Left = &TreeNode{}
					queue.Add(buildTreeInfo{
						info.in, info.pre + 1, lleft, info.node.Left,
					})
				}
				if lright > 0 {
					info.node.Right = &TreeNode{}
					queue.Add(buildTreeInfo{
						i + 1, info.pre + 1 + lleft, lright, info.node.Right,
					})
				}
				break
			}
		}
	}
	return root
}

//-------------------------

// type ListNodeAbstract struct {
// 	Val  interface{}
// 	Next *ListNodeAbstract
// }
// type Queue struct {
// 	start *ListNodeAbstract
// 	end   *ListNodeAbstract
// }

// func (q *Queue) Add(v interface{}) {
// 	if q.start == nil {
// 		q.start = &ListNodeAbstract{Val: v}
// 		q.end = q.start
// 	} else {
// 		q.end.Next = &ListNodeAbstract{Val: v}
// 		q.end = q.end.Next
// 	}
// }

// func (q *Queue) Poll() (r interface{}) {
// 	r = q.start.Val
// 	if q.start == q.end {
// 		q.start = nil
// 		q.end = nil
// 	} else {
// 		q.start = q.start.Next
// 	}
// 	return
// }

// func (q *Queue) IsEmpty() bool {
// 	return q.start == nil
// }
