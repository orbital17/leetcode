package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	i, j := 0, len(nums)
	leftR := -1
Loop:
	for {
		m := (j + i) / 2
		switch {
		case target == nums[m] && (m == 0 || nums[m-1] < target):
			leftR = m
			break Loop
		case j <= i+1:
			return []int{-1, -1}
		case nums[m] >= target:
			j = m
		default:
			i = m
		}
	}
	i, j = leftR, len(nums)
	for {
		m := (j + i) / 2
		switch {
		case target == nums[m] && (m == len(nums)-1 || nums[m+1] > target):
			return []int{leftR, m}
		case nums[m] <= target:
			i = m
		default:
			j = m
		}
	}
}
