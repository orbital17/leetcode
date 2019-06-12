package lib

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}
	if l.Next != nil {
		return fmt.Sprintf("%v -> %v", l.Val, l.Next)
	}
	return fmt.Sprint(l.Val)
}

func (l *ListNode) Reverse() *ListNode {
	var prev, next *ListNode
	for l != nil {
		next = l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}
