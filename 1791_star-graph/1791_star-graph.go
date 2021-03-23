package leetcode

func findCenter(edges [][]int) int {
	a, b := edges[0], edges[1]
	if a[0] == b[0] || a[0] == b[1] {
		return a[0]
	}
	return a[1]
}
