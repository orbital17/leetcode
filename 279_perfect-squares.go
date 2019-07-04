package leetcode

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i*i < n+1; i++ {
		dp[i*i] = 1
	}
	for i := 2; i < n+1; i++ {
		if dp[i] > 0 {
			continue
		}
		min := dp[i-1]
		for j := 1; i-j*j > 0; j++ {
			if dp[i-j*j] < min {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}
