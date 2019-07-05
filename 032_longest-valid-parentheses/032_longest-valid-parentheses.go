package leetcode

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	max := 0
	var prev int
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					prev = dp[i-2]
				} else {
					prev = 0
				}
				dp[i] = prev + 2
			}
			if dp[i-1] > 0 && i >= dp[i-1]+1 && s[i-1-dp[i-1]] == '(' {
				prev = dp[i-1] + 2
				if i-1-dp[i-1] > 0 {
					prev += dp[i-1-dp[i-1]-1]
				}
				if prev > dp[i] {
					dp[i] = prev
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}
