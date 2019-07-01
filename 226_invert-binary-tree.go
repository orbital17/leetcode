package leetcode

import . "./lib"

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	stack := &Stack{}
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*TreeNode)
		if cur == nil {
			continue
		}
		cur.Left, cur.Right = cur.Right, cur.Left
		stack.Push(cur.Right)
		stack.Push(cur.Left)
	}
	return root
}

// type ListNodeAbstract struct {
// 	Val  interface{}
// 	Next *ListNodeAbstract
// }

// type Stack struct {
// 	top *ListNodeAbstract
// }

// func (s *Stack) Push(v interface{}) {
// 	s.top = &ListNodeAbstract{
// 		Val:  v,
// 		Next: s.top,
// 	}
// }

// func (s *Stack) Pop() (r interface{}) {
// 	r = s.top.Val
// 	s.top = s.top.Next
// 	return
// }

// func (s *Stack) IsEmpty() bool {
// 	return s.top == nil
// }

// func (s *Stack) Top() interface{} {
// 	return s.top.Val
// }
