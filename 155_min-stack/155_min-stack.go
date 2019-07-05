package leetcode

import . "../lib"

type item155 struct {
	val int
	min int
}

type MinStack struct {
	stack *Stack
}

// func Constructor() MinStack {
// 	return MinStack{
// 		&Stack{},
// 	}
// }

func (this *MinStack) Push(x int) {
	item := item155{val: x}
	if this.stack.IsEmpty() {
		item.min = x
	} else {
		item.min = this.stack.Top().(item155).min
		if x < item.min {
			item.min = x
		}
	}
	this.stack.Push(item)
}

func (this *MinStack) Pop() {
	this.stack.Pop()
}

func (this *MinStack) Top() int {
	return this.stack.Top().(item155).val
}

func (this *MinStack) GetMin() int {
	return this.stack.Top().(item155).min
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
