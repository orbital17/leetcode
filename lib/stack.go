package lib

type Stack struct {
	top *ListNode
}

func (s *Stack) Push(v int) {
	s.top = &ListNode{
		Val:  v,
		Next: s.top,
	}
}

func (s *Stack) Pop() (r int) {
	r = s.top.Val
	s.top = s.top.Next
	return
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
