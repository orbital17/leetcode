package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)
	prefix := 0
	minPrefix := 0
	sum := nums[0]
	for i := 0; i < n; i++ {
		v := nums[i]
		prefix += v
		if prefix-minPrefix > sum {
			sum = prefix - minPrefix
		}
		if prefix < minPrefix {
			minPrefix = prefix
		}
	}
	return sum
}
