package leetcode

func numTrees(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	if n < 2 {
		return 1
	}
	for j := 2; j <= n; j++ {
		for i := 1; i <= j; i++ {
			f[j] += f[j-i] * f[i-1]
		}
	}
	return f[n]
}
