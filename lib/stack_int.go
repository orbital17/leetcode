package lib

type StackInt struct {
	top *ListNode
}

func (s *StackInt) Push(v int) {
	s.top = &ListNode{
		Val:  v,
		Next: s.top,
	}
}

func (s *StackInt) Pop() (r int) {
	r = s.top.Val
	s.top = s.top.Next
	return
}

func (s *StackInt) IsEmpty() bool {
	return s.top == nil
}
