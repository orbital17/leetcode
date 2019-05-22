package leetcode

func copyBoard(board [][]byte) [][]byte {
	res := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		res[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			res[i][j] = board[i][j]
		}
	}
	return res
}

func solveSudoku(board [][]byte) {
	_ = solveNext(board, 0, 0)
}

func solveNext(board [][]byte, row, col int) bool {
	if board[row][col] != byte('.') {
		if row == 8 && col == 8 {
			return true
		}
		return solveNext(board, (row+1)%9, col+(row+1)/9)
	}
	for i := 0; i < 9; i++ {
		board[row][col] = byte(int('1') + i)
		if isValid(board, row, col) {
			if solveNext(board, row, col) {
				return true
			}
		}
	}
	board[row][col] = byte('.')
	return false
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
