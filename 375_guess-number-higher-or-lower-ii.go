package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	//d[i][j] is result for [i+1, j+1) interval
	//the goal is d[0][n]
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l
			r := 1<<31 - 1 //max int32
			for k := i; k < j; k++ {
				rk := k + 1 + max(dp[i][k], dp[k+1][j])
				if rk < r {
					r = rk
				}
			}
			dp[i][j] = r
		}
	}
	return dp[0][n]
}
