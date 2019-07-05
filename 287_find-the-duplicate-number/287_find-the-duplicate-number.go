package leetcode

func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
		if fast == slow {
			break
		}
	}
	fast = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}
