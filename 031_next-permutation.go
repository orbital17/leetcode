package leetcode

func reversePerm(nums []int, i, j int) {
	for j > i {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 2
	for {
		if i < 0 {
			reversePerm(nums, 0, len(nums)-1)
			return
		}
		if nums[i+1] > nums[i] {
			break
		}
		i--
	}
	j := len(nums) - 1
	for nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	if len(nums)-i > 2 {
		reversePerm(nums, i+1, len(nums)-1)
	}
	return
}
