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
