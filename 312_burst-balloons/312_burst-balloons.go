package leetcode

func maxCoins(nums []int) int {
	b := make([]int, 1, len(nums)+2)
	b[0] = 1
	for _, v := range nums {
		if v > 0 {
			b = append(b, v)
		}
	}
	b = append(b, 1)
	n := len(b)
	if n == 2 {
		return 0
	}
	if n == 3 {
		return b[1]
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for k := 2; k < n; k++ {
		for left := 0; left+k < n; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				sum := b[i] * b[left] * b[right]
				sum += dp[left][i] + dp[i][right]
				if sum > dp[left][right] {
					dp[left][right] = sum
				}
			}
		}
	}
	return dp[0][n-1]
}
