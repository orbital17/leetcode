package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for {
		switch {
		case target == matrix[i][j]:
			return true
		case target < matrix[i][j] && j > 0:
			j--
		case target > matrix[i][j] && i < n-1:
			i++
		default:
			return false
		}
	}
	panic("")
}
