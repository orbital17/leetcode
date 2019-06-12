package lib

type TreeNodeList struct {
	Val  *TreeNode
	Next *TreeNodeList
}

type TreeStack struct {
	top *TreeNodeList
}

func (s *TreeStack) Push(v *TreeNode) {
	s.top = &TreeNodeList{
		Val:  v,
		Next: s.top,
	}
}

func (s *TreeStack) Pop() (r *TreeNode) {
	r = s.top.Val
	s.top = s.top.Next
	return
}

func (s *TreeStack) IsEmpty() bool {
	return s.top == nil
}
