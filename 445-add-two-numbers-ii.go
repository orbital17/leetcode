package leetcode

import "fmt"

//ListNode - the node in the linked list
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1, r2 := reverse(l1), reverse(l2)
	var prev *ListNode
	var sum int
	for r1 != nil || r2 != nil || sum > 9 {
		sum = sum / 10
		if r1 != nil {
			sum += r1.Val
			r1 = r1.Next
		}
		if r2 != nil {
			sum += r2.Val
			r2 = r2.Next
		}
		prev = &ListNode{sum % 10, prev}
	}
	return prev
}

func reverse(l *ListNode) *ListNode {
	var prev, next *ListNode
	for l != nil {
		next = l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}
