package leetcode

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		v := nums[i]
		if i+v > max {
			max = i + v
			if max >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
