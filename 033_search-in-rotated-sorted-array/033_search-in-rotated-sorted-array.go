package leetcode

func search033(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	i, j := 0, len(nums)
	for j > i+1 {
		m := (j + i) / 2
		if nums[m] == target {
			return m
		}
		pivotOnTheRight := nums[m] > nums[i]
		switch {
		case pivotOnTheRight && nums[m] > target && nums[i] <= target:
			j = m
		case !pivotOnTheRight && (nums[m] > target || nums[i] <= target):
			j = m
		default:
			i = m
		}
	}
	if nums[i] == target {
		return i
	}
	return -1
}
