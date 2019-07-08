package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}
	return res
}
