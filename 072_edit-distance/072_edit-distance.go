package leetcode

import . "../lib"

func minDistance(a string, b string) int {
	if len(b) > len(a) {
		a, b = b, a
	}
	if len(b) == 0 {
		return len(a)
	}
	n, m := len(a), len(b)
	dp := make([]int, m+1)
	for j := 0; j < m+1; j++ {
		dp[j] = j
	}
	for i := 1; i < n+1; i++ {
		prev := dp[0]
		dp[0] = i
		for j := 1; j < m+1; j++ {
			cur := dp[j]
			if a[i-1] == b[j-1] {
				dp[j] = prev
			} else {
				dp[j] = Min(prev, Min(cur, dp[j-1])) + 1
			}
			prev = cur
		}
	}
	return dp[m]
}
