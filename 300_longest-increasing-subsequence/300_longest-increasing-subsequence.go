package leetcode

func binarySearch(a []int, v int) int {
	l, r := 0, len(a)
	if r == 0 || r == 1 && a[0] >= v {
		return 0
	}
	for r > l+1 {
		m := (l + r) / 2
		if a[m-1] >= v {
			r = m
		} else {
			l = m
		}
		if a[l] >= v {
			return l
		}
	}
	return r
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, 0, n)
	dp = append(dp, nums[0])
	for i := 0; i < n; i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			r := binarySearch(dp, nums[i])
			dp[r] = nums[i]
		}
	}
	return len(dp)
}

func lengthOfLISn2(nums []int) int {
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
