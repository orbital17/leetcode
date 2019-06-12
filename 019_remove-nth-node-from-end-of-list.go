package leetcode

import . "./lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	cur, beforeRemove := dummy, dummy
	dist := 0
	for cur != nil {
		cur = cur.Next
		if dist <= n {
			dist++
		} else {
			beforeRemove = beforeRemove.Next
		}
	}
	beforeRemove.Next = beforeRemove.Next.Next
	return dummy.Next

}
