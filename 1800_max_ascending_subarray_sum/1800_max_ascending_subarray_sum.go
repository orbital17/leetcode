package leetcode

func maxAscendingSum(nums []int) int {
	var cur, max int
	if len(nums) == 0 {
		return 0
	}
	cur = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur += nums[i]
		} else {
			if cur > max {
				max = cur
			}
			cur = nums[i]
		}
	}
	if cur > max {
		return cur
	}
	return max
}
