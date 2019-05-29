package leetcode

func isMatch(s string, p string) bool {

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true //empty string, empty pattern

	beforeStar := false
	for j := len(p) - 1; j >= 0; j-- {
		if p[j] == '*' {
			beforeStar = true
			continue
		}
		if beforeStar && dp[len(s)][j+2] {
			dp[len(s)][j] = true //empty string, only stars in pattern
		}
		for i := len(s) - 1; i >= 0; i-- {
			if beforeStar {
				dp[i][j] =
					dp[i][j+2] ||
						dp[i+1][j] && (p[j] == '.' || s[i] == p[j])
			} else {
				dp[i][j] = dp[i+1][j+1] && (p[j] == '.' || s[i] == p[j])
			}
		}
		if beforeStar {
			beforeStar = false
		}
	}

	return dp[0][0]
}

/*
abcde ad*e

1 1 0 0
0 1 0 0
0 1 0 0
0 1 0 0
0 1 0 1

match i j =
    p[j] == '.' || s[i] == p[i]

dp i j =
    if *
        0
    if before *
        d i j+2 || d i+1 j && match pj si
    if pj is symbol
        d i+1 j+1 && match pj si


ab .*c

0 0 0| 0
0 0 0  0
-
0 0 0  1


aa a*
0 0|0
0 0 0
-
1 1 1
*/
