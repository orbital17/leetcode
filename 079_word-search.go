package leetcode

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if wordSearch(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordSearch(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}
	if word[k] != board[i][j] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	save := board[i][j]
	board[i][j] = 0
	result := wordSearch(board, word, i+1, j, k+1) ||
		wordSearch(board, word, i, j+1, k+1) ||
		wordSearch(board, word, i-1, j, k+1) ||
		wordSearch(board, word, i, j-1, k+1)
	board[i][j] = save
	return result
}
