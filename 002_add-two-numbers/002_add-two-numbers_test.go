package leetcode

import "fmt"
import . "../lib"

func Example_addTwoNumbers() {
	a := &ListNode{7,
		&ListNode{2,
			&ListNode{4,
				&ListNode{3, nil},
			},
		},
	}
	b := &ListNode{5,
		&ListNode{6,
			&ListNode{4,
				nil,
			},
		},
	}
	fmt.Print(addTwoNumbers(a, b))
	//Output:
	//2 -> 9 -> 8 -> 3
}
