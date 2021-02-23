package leetcode

import (
	"math"
)

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
