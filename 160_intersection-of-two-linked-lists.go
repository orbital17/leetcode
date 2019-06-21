package leetcode

import . "./lib"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	looped := false
	for a != b {
		a = a.Next
		if a == nil {
			if looped {
				return nil
			}
			a = headB
			looped = true
		}
		b = b.Next
		if b == nil {
			b = headA
		}
	}
	return a
}
