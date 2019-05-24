package leetcode

func PredictTheWinner(nums []int) bool {
	return predictTheWinner(nums) >= 0
}

func predictTheWinner(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
	}
	for s := len(nums); s >= 0; s-- {
		for e := s + 1; e < len(nums); e++ {
			a := nums[s] - dp[e]
			b := nums[e] - dp[e-1]
			if a > b {
				dp[e] = a
			} else {
				dp[e] = b
			}
		}
	}
	return dp[len(nums)-1]
}

// O(2^n)
// f(n) == 2 * f(n-1)
func predictTheWinnerScoresRecursive(nums []int, i, j int) (int, int) {
	if i == j {
		return 0, 0
	}
	b1, a1 := predictTheWinnerScoresRecursive(nums, i+1, j)
	b2, a2 := predictTheWinnerScoresRecursive(nums, i, j-1)
	a1 += nums[i]
	a2 += nums[j-1]
	if a1 > a2 {
		return a1, b1
	}
	return a2, b2
}

// O(n^2)
func predictTheWinnerScores(nums []int) (int, int) {
	var i, j, l, k uint16
	var a1, b1, a2, b2 int
	r1 := make(map[uint16]int)
	r2 := make(map[uint16]int)
	lenNums := uint16(len(nums))
	for i = 0; i <= lenNums; i++ {
		k = i<<8 | i
		r1[k], r2[k] = 0, 0
	}
	for l = 1; l <= lenNums; l++ {
		for i = 0; i+l <= lenNums; i++ {
			j = i + l
			k = (i+1)<<8 | j
			b1, a1 = r1[k], r2[k]
			k = i<<8 | (j - 1)
			b2, a2 = r1[k], r2[k]
			a1 += nums[i]
			a2 += nums[j-1]
			k = i<<8 | j
			if a1 > a2 {
				r1[k], r2[k] = a1, b1
			} else {
				r1[k], r2[k] = a2, b2
			}
		}
	}
	i, j = 0, lenNums
	k = i<<8 | j
	return r1[k], r2[k]
}
