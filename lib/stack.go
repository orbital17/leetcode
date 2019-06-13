package lib

type ListNodeAbstract struct {
	Val  interface{}
	Next *ListNodeAbstract
}

type Stack struct {
	top *ListNodeAbstract
}

func (s *Stack) Push(v interface{}) {
	s.top = &ListNodeAbstract{
		Val:  v,
		Next: s.top,
	}
}

func (s *Stack) Pop() (r interface{}) {
	r = s.top.Val
	s.top = s.top.Next
	return
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
