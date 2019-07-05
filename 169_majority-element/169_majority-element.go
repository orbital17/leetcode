package leetcode

// import . "../lib"

func majorityElement(nums []int) int {
	var i, count, cur int
	for i = 0; i < len(nums); i++ {
		if count == 0 {
			count++
			cur = nums[i]
		} else if nums[i] == cur {
			count++
		} else {
			count--
		}
	}
	return cur
}
