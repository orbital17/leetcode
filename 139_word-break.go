package leetcode

func wordBreak(s string, wordDict []string) bool {

	set := make(map[string]bool)
	for _, v := range wordDict {
		set[v] = true
	}
	exists := func(v string) bool {
		_, ok := set[v]
		return ok
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && exists(s[j:i]) {
				dp[i] = true
			}
		}
	}
	return dp[n]
}
