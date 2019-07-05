package leetcode

import . "../lib"

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
