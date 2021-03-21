package leetcode

import (
	"fmt"

	. "github.com/orbital17/leetcode/lib"
)

func Example_mergeKLists() {
	lists := []*ListNode{
		NewListNode(1, 4, 5),
	}
	actual := mergeKLists(lists)
	fmt.Print(actual)
	//Output:
	//1 -> 4 -> 5
}

func Example_mergeKLists2() {
	lists := []*ListNode{
		NewListNode(1, 4, 5),
		NewListNode(1, 3, 4),
		NewListNode(2, 6),
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
