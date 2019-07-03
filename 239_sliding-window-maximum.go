package leetcode

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	dequeue := list.New()
	res := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		for dequeue.Len() != 0 &&
			dequeue.Front().Value.(int) < i-k+1 {
			dequeue.Remove(dequeue.Front())
		}
		for dequeue.Len() != 0 &&
			nums[dequeue.Back().Value.(int)] < nums[i] {
			dequeue.Remove(dequeue.Back())
		}

		dequeue.PushBack(i)
		if i+1 >= k {
			res[i-k+1] = nums[dequeue.Front().Value.(int)]
		}
	}
	return res
}
