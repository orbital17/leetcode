package leetcode

import "math/rand"

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)
	for {
		px := rand.Intn(r-l) + l
		nums[l], nums[px] = nums[px], nums[l]
		px = l
		p := nums[px]
		for i := l + 1; i < r; i++ {
			if nums[i] >= p {
				nums[i], nums[px+1] = nums[px+1], nums[i]
				nums[px], nums[px+1] = nums[px+1], nums[px]
				px++
			}
		}
		if px == k-1 {
			return nums[px]
		} else if px < k-1 {
			l = px
		} else {
			r = px
		}
	}
}
