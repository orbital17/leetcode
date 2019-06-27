package leetcode

import . "./lib"

func reverseList_(l *ListNode) *ListNode {
	var prev, next *ListNode
	for l != nil {
		next = l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}

func reverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	n := l.Next
	l.Next = nil
	r := reverseList(n)
	n.Next = l
	return r
}
