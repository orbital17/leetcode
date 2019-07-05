package leetcode

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	p := partition(nums)
	sortArray(nums[:p])
	sortArray(nums[p+1:])
	return nums
}

func partition(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	p := nums[n-1]
	i := 0
	for j := 0; j < n; j++ {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[n-1] = nums[n-1], nums[i]
	return i
}
