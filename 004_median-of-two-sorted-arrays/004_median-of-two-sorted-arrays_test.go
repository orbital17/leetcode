package leetcode

import "fmt"

func Example_findMedianSortedArrays() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5}
	res := findMedianSortedArrays(a1, a2)
	fmt.Print(res)
	//Output:
	// 3
}
