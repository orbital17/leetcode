package leetcode

import "fmt"
import "math"

func ExampleAdd() {
	a := &ListNode{7,
		&ListNode{2,
			&ListNode{4,
				&ListNode{3,nil},
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
	//7 -> 8 -> 0 -> 7
}

func f(n int, bits float64) float64 {
	e := math.Pow10(n) / math.Exp2(bits/2)
	e = e * e / 2
	return 1 - math.Exp(-e)
}

func ExampleMath() {
	r := f(8, 72)
	fmt.Print(r)
	//Output:
	//
}