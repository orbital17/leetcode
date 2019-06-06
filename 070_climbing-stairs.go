package leetcode

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
