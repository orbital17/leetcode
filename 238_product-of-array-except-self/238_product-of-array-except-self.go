package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	cur := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = cur
		cur *= nums[i]
	}
	cur = 1
	for i := 0; i < n; i++ {
		res[i] = res[i] * cur
		cur *= nums[i]
	}
	return res
}
