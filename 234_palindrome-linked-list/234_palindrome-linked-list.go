package leetcode

import . "github.com/orbital17/leetcode/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	cur := head
	n := 0
	for cur != nil {
		n++
		cur = cur.Next
	}
	m := (n + 1) / 2
	cur = head
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	var prev *ListNode
	for cur != nil {
		n := cur.Next
		cur.Next = prev
		prev = cur
		cur = n
	}
	a := prev
	b := head
	for a != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}
