package leetcode

func copyPars(parsOpen []string, parsClose []string) []string {
	result := make([]string, len(parsOpen)+len(parsClose))
	for i, v := range parsOpen {
		result[i] = v + "("
	}
	for i, v := range parsClose {
		result[i+len(parsOpen)] = v + ")"
	}
	return result
}

func generateParenthesisDP(n int) []string {
	dp := [][]string{{""}, {}}
	var open, close []string
	for i := 1; i <= 2*n; i++ {
		newDp := make([][]string, i/2+1)
		for j := 0; j <= i/2; j++ {
			if i%2 == 0 {
				if j < len(dp) {
					close = dp[j]
				} else {
					close = []string{}
				}
				if j > 0 {
					open = dp[j-1]
				} else {
					open = []string{}
				}
			} else {
				open = dp[j]
				if j+1 < len(dp) {
					close = dp[j+1]
				} else {
					close = []string{}
				}
			}
			newDp[j] = copyPars(open, close)
		}
		dp = newDp
	}
	return dp[0]
}

//------------ backtrack ---------------

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, []byte{}, 0, 2*n)
	return result
}

func backtrack(result *[]string, symbols []byte, open, n int) {
	if n == 0 {
		*result = append(*result, string(symbols))
		return
	}
	if open != 0 {
		backtrack(result, append(symbols, ')'), open-1, n-1)
	}
	if open != n {
		backtrack(result, append(symbols, '('), open+1, n-1)
	}
}
