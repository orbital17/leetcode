package leetcode

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	result := make([][]int, 0, 1<<uint(len(nums)))
	subsetsBacktrack(&result, nums, []int{}, 0)
	return result
}

func subsetsBacktrack(result *[][]int, nums []int, current []int, i int) {
	new := make([]int, len(current)+1)
	copy(new, current)
	new[len(current)] = nums[i]
	if i+1 == len(nums) {
		*result = append(*result, new)
		*result = append(*result, current)
	} else {
		subsetsBacktrack(result, nums, current, i+1)
		subsetsBacktrack(result, nums, new, i+1)
	}
}
