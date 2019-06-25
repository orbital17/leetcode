package leetcode

import . "./lib"

func rob(nums []int) int {
	robbed, notRobbed := 0, 0
	for _, v := range nums {
		robbed, notRobbed = Max(notRobbed+v, robbed), Max(robbed, notRobbed)
	}
	return Max(robbed, notRobbed)
}

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
