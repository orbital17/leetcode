package leetcode

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		for i := 1; i < len(p); i += 2 {
			if p[i] != '*' {
				return false
			}
		}
		return len(p)%2 == 0
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(p))
	}
	beforeStar := false
	for j := len(p) - 1; j >= 0; j-- {
		if p[j] == '*' {
			beforeStar = true
			continue
		}
		for i := len(s) - 1; i >= 0; i-- {
			if beforeStar {
				dp[i][j] =
					(j+2 < len(p) &&
						dp[i][j+2]) ||
						(i+1 >= len(s) ||
							dp[i+1][j]) &&
							(p[j] == '.' || s[i] == p[j])
			} else {
				if i+1 == len(s) || j+1 == len(p) {
					dp[i][j] = i+1 == len(s) && j+1 == len(p) &&
						(p[j] == '.' || s[i] == p[j])
				} else {
					dp[i][j] = dp[i+1][j+1] && (p[j] == '.' || s[i] == p[j])
				}
			}
		}
		if beforeStar {
			beforeStar = false
		}
	}

	return dp[0][0]
}
