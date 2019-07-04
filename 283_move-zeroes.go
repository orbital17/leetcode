package leetcode

func moveZeroes(nums []int) {
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			count++
		} else if count != 0 {
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
}
