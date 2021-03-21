package leetcode

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findMax(nums []int, dp []int, s int, mask int) int {
	if s >= len(nums)/2 {
		return 0
	}
	if dp[mask] > 0 {
		return dp[mask]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			newMask := 1<<i + 1<<j
			if newMask&mask == 0 {
				score := findMax(nums, dp, s+1, mask|newMask) + (s+1)*gcd(nums[i], nums[j])
				if score > dp[mask] {
					dp[mask] = score
				}
			}
		}
	}
	return dp[mask]
}

func maxScore(nums []int) int {
	dp := make([]int, 1<<len(nums))
	return findMax(nums, dp, 0, 0)
}
