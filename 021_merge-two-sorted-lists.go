package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	newEnd := dummy
	var min *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			min = l1
			l1 = l1.Next
		} else {
			min = l2
			l2 = l2.Next
		}
		newEnd.Next = min
		newEnd = min
	}
	if l1 == nil {
		newEnd.Next = l2
	} else {
		newEnd.Next = l1
	}
	return dummy.Next
}
