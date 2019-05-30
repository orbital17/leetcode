package leetcode

func twoSumN2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)
	for i, v := range nums {
		i0, ok := indices[target-v]
		if ok {
			return []int{i0, i}
		}
		indices[v] = i
	}
	return nil
}
