package leetcode

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max, s := 0, 0
	for r > l {
		if height[r] > height[l] {
			s = (r - l) * height[l]
			l++
		} else {
			s = (r - l) * height[r]
			r--
		}
		if s > max {
			max = s
		}
	}
	return max
}
