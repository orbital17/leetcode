package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == 1<<31-1 {
		return -1
	} else {
		return dp[amount]
	}
}
