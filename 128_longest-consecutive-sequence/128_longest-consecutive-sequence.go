package leetcode

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	exists := func(n int) bool {
		_, ok := set[n]
		return ok
	}

	max := 0
	for _, v := range nums {
		if !exists(v - 1) {
			cur := v
			curStreak := 1
			for exists(cur + 1) {
				cur++
				curStreak++
			}
			if curStreak > max {
				max = curStreak
			}
		}
	}
	return max
}
