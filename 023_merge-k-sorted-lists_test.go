package leetcode

import (
	"fmt"
)

func newListNode(nums ...int) *ListNode {
	var current *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		current = &ListNode{
			nums[i],
			current,
		}
	}
	return current
}

func Example_mergeKLists() {
	lists := []*ListNode{
		newListNode(1, 4, 5),
	}
	actual := mergeKLists(lists)
	fmt.Print(actual)
	//Output:
	//1 -> 4 -> 5
}

func Example_mergeKLists2() {
	lists := []*ListNode{
		newListNode(1, 4, 5),
		newListNode(1, 3, 4),
		newListNode(2, 6),
	}
	actual := mergeKLists(lists)
	fmt.Print(actual)
	//Output:
	//1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
}
func Example_mergeKLists3() {
	lists := []*ListNode{}
	actual := mergeKLists(lists)
	fmt.Print(actual)
	//Output:
	//<nil>
}

func Example_mergeKLists4() {
	lists := []*ListNode{nil}
	actual := mergeKLists(lists)
	fmt.Print(actual)
	//Output:
	//<nil>
}
