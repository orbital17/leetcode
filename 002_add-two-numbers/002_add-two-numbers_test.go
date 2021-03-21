package leetcode

import (
	"fmt"

	. "github.com/orbital17/leetcode/lib"
)

func Example_addTwoNumbers() {
	a := NewListNode(7, 2, 4, 3)
	b := NewListNode(5, 6, 4)
	fmt.Print(addTwoNumbers(a, b))
	//Output:
	//2 -> 9 -> 8 -> 3
}
