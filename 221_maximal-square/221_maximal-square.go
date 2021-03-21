package leetcode

import . "github.com/orbital17/leetcode/lib"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	dp := make([]int, m)
	prev := 0
	res := 0
	for i := 0; i < n; i++ {
		prev = dp[0]
		if matrix[i][0] == '1' {
			dp[0] = 1
			res = Max(res, 1)
		} else {
			dp[0] = 0
		}
		for j := 1; j < m; j++ {
			var next int
			if matrix[i][j] == '1' {
				next = Min(Min(prev, dp[j]), dp[j-1]) + 1
			} else {
				next = 0
			}
			prev, dp[j] = dp[j], next
			res = Max(res, dp[j])
		}
	}
	return res * res
}

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func Min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
