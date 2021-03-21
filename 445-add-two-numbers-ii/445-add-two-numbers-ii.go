package leetcode

import . "github.com/orbital17/leetcode/lib"

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
		prev = &ListNode{Val: sum % 10, Next: prev}
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
