package leetcode

// import . "./lib"

func majorityElement(nums []int) int {
	var i, count, cur int
	for i < len(nums) {
		cur = nums[i]
		count = 1
		i++
		for count != 0 && i < len(nums) {
			if cur == nums[i] {
				count++
			} else {
				count--
			}
			i++
		}
	}
	return cur
}
