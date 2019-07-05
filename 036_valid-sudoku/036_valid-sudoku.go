package leetcode

func isValidSudoku(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if board[j][i] != '.' && !isValid(board, j, i) {
				return false
			}
		}
	}
	return true
}

func isValid(b [][]byte, row, col int) bool {
	v := b[row][col]
	ci, ri := col-col%3, row-row%3
	var c, r int
	for i := 0; i < 9; i++ {
		if b[row][i] == v && i != col {
			return false
		}
		if b[i][col] == v && i != row {
			return false
		}
		c, r = ci+i%3, ri+i/3
		if b[r][c] == v && !(r == row && c == col) {
			return false
		}
	}
	return true
}
