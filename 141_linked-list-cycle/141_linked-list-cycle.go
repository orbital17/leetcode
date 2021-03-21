package leetcode

import . "github.com/orbital17/leetcode/lib"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycleMap(head *ListNode) bool {
	if head == nil {
		return false
	}
	set := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if _, ok := set[cur]; ok {
			return true
		}
		set[cur] = true
		cur = cur.Next
	}
	return false
}
