package leetcode

import . "github.com/orbital17/leetcode/lib"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{
		Next: head,
	}
	for step := 1; step < n; step *= 2 {
		tail := dummy
		cur := dummy.Next
		for cur != nil {
			left := cur
			right := listFirstN(cur, step)
			cur = listFirstN(right, step)
			tail = mergeList(left, right, tail)
		}
	}
	return dummy.Next
}

func listFirstN(head *ListNode, n int) (next *ListNode) {
	for head != nil && n > 1 {
		head = head.Next
		n--
	}
	if head == nil {
		return nil
	}
	next = head.Next
	head.Next = nil
	return
}

func mergeList(a, b, prev *ListNode) (last *ListNode) {
	cur := prev
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			cur = a
			a = a.Next
		} else {
			cur.Next = b
			cur = b
			b = b.Next
		}
	}
	if b != nil {
		a = b
	}
	cur.Next = a
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
