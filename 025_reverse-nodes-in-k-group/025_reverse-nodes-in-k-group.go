package leetcode

import . "github.com/orbital17/leetcode/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	if n < k {
		return head
	}
	n = n / k * k
	cur, firstInBlock := head, head
	dummy := &ListNode{}
	prev, firstInPrevBlock := dummy, dummy
	i := 0
	for cur != nil && i < n {
		if i%k == 0 && i > 0 {
			firstInPrevBlock.Next = prev
			firstInPrevBlock = firstInBlock
			firstInBlock = cur
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		i++
	}
	firstInPrevBlock.Next = prev
	firstInBlock.Next = cur
	return dummy.Next
}
