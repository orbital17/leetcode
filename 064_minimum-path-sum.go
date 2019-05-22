package leetcode

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for j := m - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if i == n-1 && j == m-1 {
				continue
			}
			if i == n-1 {
				grid[j][i] += grid[j+1][i]
			} else if j == m-1 {
				grid[j][i] += grid[j][i+1]
			} else {
				grid[j][i] += min(grid[j+1][i], grid[j][i+1])
			}
		}
	}
	return grid[0][0]
}
