package leetcode

import (
	"fmt"
	"math"

	. "github.com/orbital17/leetcode/lib"
)

func Example_addTwoNumbers() {
	a := NewListNode(7, 2, 4, 3)
	b := NewListNode(5, 6, 4)
	fmt.Print(addTwoNumbers(a, b))
	//Output:
	//7 -> 8 -> 0 -> 7
}

func f(n int, bits float64) float64 {
	e := math.Pow10(n) / math.Exp2(bits/2)
	e = e * e / 2
	return 1 - math.Exp(-e)
}

func Example_math() {
	r := f(8, 72)
	fmt.Print(r)
	//Output:
	//
}
