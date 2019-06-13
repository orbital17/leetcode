package lib

type ListNodeAbstract struct {
	Val  interface{}
	Next *ListNodeAbstract
}

type StackAbstract struct {
	top *ListNodeAbstract
}

func (s *StackAbstract) Push(v interface{}) {
	s.top = &ListNodeAbstract{
		Val:  v,
		Next: s.top,
	}
}

func (s *StackAbstract) Pop() (r interface{}) {
	r = s.top.Val
	s.top = s.top.Next
	return
}

func (s *StackAbstract) IsEmpty() bool {
	return s.top == nil
}
