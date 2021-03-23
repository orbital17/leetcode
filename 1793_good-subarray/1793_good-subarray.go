package leetcode

func maximumScore(nums []int, k int) int {
	minV := nums[k]
	i, j := k, k
	maxScore := 0
	extend := func() {
		for i-1 >= 0 && nums[i-1] >= minV {
			i--
		}
		for j+1 < len(nums) && nums[j+1] >= minV {
			j++
		}
		score := (j - i + 1) * minV
		if score > maxScore {
			maxScore = score
		}
	}
	extend()
	for i > 0 || j < len(nums)-1 {
		if i > 0 && j < len(nums)-1 {
			if nums[i-1] > nums[j+1] {
				minV = nums[i-1]
			} else {
				minV = nums[j+1]
			}
		} else if i > 0 {
			minV = nums[i-1]
		} else {
			minV = nums[j+1]
		}
		extend()
	}
	return maxScore
}
